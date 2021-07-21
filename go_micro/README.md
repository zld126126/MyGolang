# go_micro
go_micro demo

brew upgrade go

go get -d -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/google/wire/cmd/wire
go get -u github.com/micro/protoc-gen-micro

protoc --proto_path=. --micro_out=. --go_out=. *.proto

运行步骤:
pkg/main.go client.go
