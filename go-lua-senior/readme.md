# go-lua-senior
    go-lua 高级示例

## 1.注册自定义函数
    p.Lua.Register("goEvent", p.goEvent)
## 2.封装lua虚拟机
    lua.NewState()
    lua.OpenLibraries(m.Lua)
## 3.goFunc-luaFunc互相调用
    goFunc("goFunc from test1.lua")
## 4.goFunc-luaFunc回调
    cb(compute)