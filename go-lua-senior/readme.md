# go-lua-senior
    github.com/Shopify/go-lua 高级示例
  - [1.注册自定义函数](#1注册自定义函数)
  - [2.封装lua虚拟机](#2封装lua虚拟机)
  - [3.goFunc-luaFunc互相调用](#3gofunc-luafunc互相调用)
  - [4.goFunc-luaFunc回调](#4gofunc-luafunc回调)

## 1.注册自定义函数
    p.Lua.Register("goEvent", p.goEvent)
## 2.封装lua虚拟机
    lua.NewState()
    lua.OpenLibraries(m.Lua)
## 3.goFunc-luaFunc互相调用
    goFunc("goFunc from test1.lua")
## 4.goFunc-luaFunc回调
    cb(compute)