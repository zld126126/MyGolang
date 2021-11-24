# go_dll
- [go_dll](#go_dll)
  - [1.环境](#1环境)
  - [2.运行](#2运行)
  - [3.编译](#3编译)

## 1.环境
    vs2013+go1.13+x86 dll文件
    
## 2.运行
    go_dll.exe
    
## 3.编译
    vs2013:
        重新生成解决方案
        \cpp_dll\Debug\cpp_dll.dll
    
    goland:
        Mac:
            CGO_ENABLED=1 GOOS=windows GOARCH=386 go build main.go
        
        Windows:
            SET CGO_ENABLED=1
            SET GOOS=windows
            SET GOARCH=386
            go build main.go
            
## 4.生成c++类库:
     .a .h .dll .lib 文件生成
### 4.1 生成.a/.h
    go build -buildmode=c-archive go_lib.go

### 4.2 生成.dll/.lib
    gcc go_lib.def go_lib.a -shared -lwinmm -lWs2_32 -o go_lib.dll -Wl,--out-implib,go_lib.lib

