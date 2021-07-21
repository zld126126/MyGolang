# go_micro_v2
go_micro_v2_demo by dongtech

预备步骤:
brew upgrade go
GO111MODULE=on go get github.com/micro/micro/v2@v2.9.3
go get -d -u github.com/golang/protobuf/protoc-gen-go go get -u github.com/google/wire/cmd/wire go get -u github.com/micro/protoc-gen-micro
protoc --proto_path=. --micro_out=. --go_out=. *.proto

运行步骤: 
1.运行main.go
2.执行http://localhost:10086/greeter
3.访问http://localhost:10086/greeter/dong

参考资料:
https://github.com/yuedun/micro-service
