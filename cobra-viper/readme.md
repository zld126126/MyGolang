设置proxy

export GO111MODULE=on

export GOPROXY=https://goproxy.io

go get组件

go get github.com/spf13/viper

go get -u github.com/spf13/cobra/cobra

问题：

ambiguous import: found github.com/ugorji/go/codec in multiple modules

解决：

go get github.com/ugorji/go@v1.1.2

启动main.go即可

测试接口
localhost:9090/version
localhost:9090/config
