package goLua

import (
	"dongo_game_server/src/util"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
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

type LuaInterface struct{}

var Lua_Interface = &LuaInterface{}
var Lua_Instance = luaPool.Get()

const LuaPath = "/src/goLua/lua/"

func (p *LuaInterface) getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 默认lua文件路径为/src/web/scripts/lua
func (p *LuaInterface) GetLuaDirectory() string {
	return p.getCurrentDirectory() + LuaPath
}

func (p *LuaInterface) DoString(s string) {
	err := Lua_Instance.DoString(s)
	if err != nil {
		logrus.WithError(err).Println("lua do string error")
	}
}

func (p *LuaInterface) DoFile(fileName string) {
	path := p.GetLuaDirectory() + fileName
	err := Lua_Instance.DoFile(path)
	if err != nil {
		logrus.WithError(err).Println("lua do file error")
	}
}

func (p *LuaInterface) DoFileWithRes(fileName string, funcName string) int {
	// 加载fib.lua
	path := p.GetLuaDirectory() + fileName
	if err := Lua_Instance.DoFile(path); err != nil {
		panic(err)
	}
	// 调用fib(n)
	err := Lua_Instance.CallByParam(lua.P{
		Fn:      Lua_Instance.GetGlobal(funcName), // 获取fib函数引用
		NRet:    1,                                // 指定返回值数量
		Protect: true,                             // 如果出现异常，是panic还是返回err
	}, lua.LNumber(10)) // 传递输入参数n=10
	if err != nil {
		util.Chk(err)
	}
	// 获取返回结果
	ret := Lua_Instance.Get(-1)
	// 从堆栈中扔掉返回结果
	Lua_Instance.Pop(1)
	// 打印结果
	res, ok := ret.(lua.LNumber)
	if ok {
		fmt.Println(int(res))
	} else {
		fmt.Println("unexpected result")
	}
	return int(res)
}

func (p *LuaInterface) Example() {
	p.DoString(`print("hello")`)
	res := p.DoFileWithRes("fib.lua", "fib")
	logrus.WithField("dofile Result", res)
}
