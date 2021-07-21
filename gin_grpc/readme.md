****grpc-test****

安装protobuf

1.brew install protobuf

2.protoc --version

安装grpc

1.go get google.golang.org/grpc

2.go get github.com/golang/protobuf/protoc-gen-go

生成 pb.go

1.protoc --go_out=plugins=grpc:. *.proto

安装gin

1.go get github.com/gin-gonic/gin

测试运行

1.serve+client

2.serve+main

****访问localhost:9090/grpc****

设置proxy

export GO111MODULE=on

export GOPROXY=https://goproxy.io
