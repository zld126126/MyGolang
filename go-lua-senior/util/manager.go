package util

import (
	"fmt"
	"github.com/Shopify/go-lua"
)

type Manager struct {
	Lua          *lua.State
	EventManager *EventManager
}

func (p *Manager) goEvent(l *lua.State) int {
	fmt.Println("goEvent")
	s, ok := l.ToString(1)
	if !ok || s == "" {
		return 0
	}

	if p.EventManager != nil {
		p.EventManager.DispatchEvent("_goEvent", s)
	}
	return 0
}

func (p *Manager) GoEvent(i interface{}) error {
	s, ok := i.(string)
	if !ok || s == "" {
		return nil
	}

	fmt.Println("_goEvent,in:", s)
	return nil
}

func (p *Manager) cb(l *lua.State) int {
	fmt.Println("cb")
	if l.IsFunction(1) {
		// 参数1 1
		l.PushInteger(1)
		// 参数2 2
		l.PushInteger(2)
		l.Call(2, 1)
		// 结果 1+2 = 3
		i, ok := l.ToInteger(1)
		if ok && i > 0 {
			fmt.Println(i)
		}
	}
	return 0
}

// 注册事件
func (p *Manager) Register() {
	p.Lua.Register("goEvent", p.goEvent)
	p.Lua.Register("cb", p.cb)
}

func NewManager() *Manager {
	m := &Manager{
		EventManager: NewEventManager(),
		Lua:          lua.NewState(),
	}

	lua.OpenLibraries(m.Lua)
	m.EventManager.RegisterEvent("_goEvent", m.GoEvent)
	m.Register()
	return m
}
