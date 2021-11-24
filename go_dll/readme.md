# go_dll
- [go_dll](#go_dll)
  - [1. 安装 && 编译 && 运行](#1-安装--编译--运行)
    - [1.1 环境](#11-环境)
    - [1.2 运行](#12-运行)
    - [1.3 编译](#13-编译)
    - [1.4 go调用c++类库](#14-go调用c类库)
  - [2. 生成c++类库:](#2-生成c类库)
    - [2.1 生成.a/.h](#21-生成ah)
    - [2.2 生成.dll/.lib](#22-生成dlllib)
    - [2.3 调用go生成的类库](#23-调用go生成的类库)

## 1. 安装 && 编译 && 运行
### 1.1 环境
    vs2013+go1.13+x86 dll文件
    
### 1.2 运行
    go_dll.exe
    
### 1.3 编译
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

### 1.4 go调用c++类库
    // 三种调用方式
	res1 := GoCallDll1(4, 5)
	fmt.Println("r1:", (int)(res1))

	res2 := GoCallDll2(3, 6)
	fmt.Println("r2:", Common(res2))

	res3 := GoCallDll3(2, 7)
	fmt.Println("r3:", (int)(res3))            
            
## 2. 生成c++类库:
     .a .h .dll .lib 文件生成
### 2.1 生成.a/.h
    go build -buildmode=c-archive go_lib.go

### 2.2 生成.dll/.lib
    gcc go_lib.def go_lib.a -shared -lwinmm -lWs2_32 -o go_lib.dll -Wl,--out-implib,go_lib.lib

### 2.3 调用go生成的类库
    fmt.Println("go call goDLL")
	GoCallDll(3, 5)

    func GoCallDll(a, b int) uintptr {
        DllTestDef := syscall.MustLoadDLL(goDLL)
        add := DllTestDef.MustFindProc("Add")

        fmt.Println("+++++++MustFindProc：", add, "+++++++")
        ret, _, err := add.Call(IntPtr(a), IntPtr(b))
        if err != nil && IsFinishError(err) {
            fmt.Println(goDLL+fmt.Sprintf(":%d+%d", a, b)+"结果为:", ret)
        } else {
            fmt.Println(fmt.Sprintf("%+v", err))
        }
        return ret
    }

