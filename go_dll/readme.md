# go_dll
## 环境
    vs2013+go1.13
    x86dll文件
    
## 运行
    go_dll.exe
    
## 编译
    vs2013:
        重新生成解决方案
    
    goland:
        Mac:
            CGO_ENABLED=1 GOOS=windows GOARCH=386 go build main.go
        
        Windows:
            SET CGO_ENABLED=1
            SET GOOS=windows
            SET GOARCH=386
            go build main.go