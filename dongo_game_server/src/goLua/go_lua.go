package goLua

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"github.com/zld126126/dongo_utils"
)

/**
参考文章:
	https://www.136.la/jingpin/show-121333.html
	https://github.com/yuin/gopher-lua

TODO:
	1.Creating a module by Go
	2.Calling Lua from Go
	3.Working with coroutines
	...

Lua类型 <--> Go类型
	LNilType -- nil
	LBool -- bool
	LNumber -- float64
	LString -- string
	LFunction
	LUserData
	LState
	LTable
	LChannel
**/

type lStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *lStatePool) New() *lua.LState {
	//L := lua.NewState()
	L := lua.NewState(lua.Options{
		CallStackSize:       120,
		MinimizeStackMemory: true,
	})
	// setting the L up here.
	// load scripts, set global variables, share channels, etc...
	return L
}

func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

func (pl *lStatePool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

// Global LState pool
var luaPool = &lStatePool{saved: make([]*lua.LState, 0, 4)}

type _LuaObject struct{}

var LuaObject = &_LuaObject{}
var Lua = luaPool.Get()

const LuaPath = "/src/goLua/lua/"

func (p *_LuaObject) getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`lua getCurrentDirectory error`)
		dongo_utils.Chk(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 默认lua文件路径为/src/web/scripts/lua
func (p *_LuaObject) GetLuaDirectory() string {
	return p.getCurrentDirectory() + LuaPath
}

func (p *_LuaObject) DoString(s string) {
	err := Lua.DoString(s)
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln("lua DoString error")
		dongo_utils.Chk(err)
	}
}

func (p *_LuaObject) DoFile(fileName string) {
	path := p.GetLuaDirectory() + fileName
	err := Lua.DoFile(path)
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln("lua DoFile error")
		dongo_utils.Chk(err)
	}
}

type LuaBaseTp int32

const (
	LuaBaseTpUnknow LuaBaseTp = iota
	LuaBaseTpNil
	LuaBaseTpBool
	LuaBaseTpNumber
	LuaBaseTpString
	LuaBaseTpFunction
	LuaBaseTpUserData
	LuaBaseTpState
	LuaBaseTpTable
	LuaBaseTpChannel
)

type LuaBaseResponse struct {
	Tp          LuaBaseTp   `json:"tp"`
	StringValue string      `json:"stringValue"`
	IntValue    int         `json:"intValue"`
	BoolValue   bool        `json:"boolValue"`
	Value       interface{} `json:"value"` // TODO 后期支持更多类型
}

func (p *_LuaObject) DoFileWithRes(fileName string, funcName string) *LuaBaseResponse {
	// 加载fib.lua
	path := p.GetLuaDirectory() + fileName
	if err := Lua.DoFile(path); err != nil {
		panic(err)
	}

	// 调用fib(n)
	err := Lua.CallByParam(lua.P{
		Fn:      Lua.GetGlobal(funcName), // 获取fib函数引用
		NRet:    1,                       // 指定返回值数量
		Protect: true,                    // 如果出现异常，是panic还是返回err
	}, lua.LNumber(10)) // 传递输入参数n=10
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln("lua DoFile error")
		return nil
	}

	// 获取返回结果
	val := Lua.Get(-1)
	var resp LuaBaseResponse

	switch val.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		res := val.(lua.LBool)
		resp = LuaBaseResponse{
			Tp:        LuaBaseTpBool,
			BoolValue: bool(res),
		}
	case lua.LTNumber:
		res := val.(lua.LNumber)
		resp = LuaBaseResponse{
			Tp:       LuaBaseTpNumber,
			IntValue: int(res),
		}
	case lua.LTString:
		res := val.(lua.LString)
		resp = LuaBaseResponse{
			Tp:          LuaBaseTpString,
			StringValue: string(res),
		}
	default:
		logrus.WithField("err", fmt.Sprintf("%+v", errors.New("convert lua unknown"))).Errorln("lua DoFileWithRes unknown error")
		resp = LuaBaseResponse{
			Tp:    LuaBaseTpUnknow,
			Value: val,
		}
	}

	// 从堆栈中扔掉返回结果
	Lua.Pop(1)
	// 打印结果
	return &resp
}

func (p *_LuaObject) double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1
}

func (p *_LuaObject) Example_Lua2Go() {
	Lua.SetGlobal("double", Lua.NewFunction(p.double))
	Lua.DoString("print(double(20))")
}

func (p *_LuaObject) Example_Go2Lua() {
	p.DoString(`print("hello")`)

	res := p.DoFileWithRes("fib.lua", "fib")
	_json, err := dongo_utils.ToJsonString(res)
	if err != nil {
		dongo_utils.Chk(err)
	}
	logrus.WithField("Tp", res.Tp).WithField("Object", _json).Println("DoFileWithRes success")
}
