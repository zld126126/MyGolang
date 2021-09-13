# dongo_game_server 开发计划
    安装protobuf
        1.brew install protobuf
        2.protoc --version

    安装grpc
        1.go get google.golang.org/grpc
        2.go get github.com/golang/protobuf/protoc-gen-go

    生成 pb.go
        1.protoc --go_out=plugins=grpc:. *.proto

## 功能:
    1 网页系统 
        1.1 用户登录管理系统
        1.2 项目系统
        1.3 数据统计系统
        1.4 Api对接文档 √
    
    2 客户端系统 
        2.1 采集数据
        2.2 socket/http/rpc 连接 √
        2.3 静态资源获取
        2.4 Lua 脚本 √
        
    3 支持系统
        3.1 定时邮件任务
        

## 问题:
    Type definition of type ‘*ast.InterfaceType’ is not supported yet. Using ‘object’ instead.
    
    go get -u github.com/swaggo/swag/cmd/swag@v1.6.7