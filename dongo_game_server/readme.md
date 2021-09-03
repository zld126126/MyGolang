# dongo_game_server 开发计划
    安装所需类库:
        go get 

        github.com/BurntSushi/toml
        github.com/dchest/captcha
        github.com/deckarep/golang-set
        github.com/dgrijalva/jwt-go
        github.com/gin-gonic/gin
        github.com/go-gomail/gomail
        github.com/golang/protobuf
        github.com/google/wire
        github.com/jinzhu/copier
        github.com/jinzhu/gorm
        github.com/pkg/errors
        github.com/robfig/cron
        github.com/sirupsen/logrus
        github.com/spf13/viper
        golang.org/x/net v0.0.0-20210825183410-e898025ed96a
        google.golang.org/grpc
        google.golang.org/protobuf
        gopkg.in/alexcesaro/quotedprintable.v3
    
    安装protobuf
        1.brew install protobuf
        2.protoc --version

    安装grpc
        1.go get google.golang.org/grpc
        2.go get github.com/golang/protobuf/protoc-gen-go

    生成 pb.go
        1.protoc --go_out=plugins=grpc:. *.proto

## 功能:
    1 用户系统
    2 聊天系统
    3 资源系统
    4 token鉴权系统
    5 socket http proto服务系统
    6 邮件系统
    7 打点系统
