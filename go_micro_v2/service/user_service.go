package service

import (
	"context"
	"strconv"

	"github.com/micro/go-micro/v2/logger"

	user "go_micro_v2/proto/user"
)

type User struct{}

// 实现了user.pb.micro.go中的UserHandler接口
func (e *User) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	logger.Info("Received QueryUserByName request:", req.GetUserName())
	// TODO 改成 数据库查询
	// rsp.User.Name = "Hello " + req.UserName//rsp.User是零值（nil），不能直接对其属性赋值，所以需要创建新对象赋值
	ID64, _ := strconv.ParseInt(req.UserID, 10, 64)
	rsp.User = &user.User{
		Id:   ID64,
		Name: req.UserName,
		Pwd:  req.UserPwd,
	}
	rsp.Success = true
	return nil
}
