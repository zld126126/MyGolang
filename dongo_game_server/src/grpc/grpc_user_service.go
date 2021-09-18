package grpc

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/database"
	"dongo_game_server/src/model"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/zld126126/dongo_utils"
	"golang.org/x/net/context"
)

type RpcUserService struct {
	DB *database.DB
}

func (p *RpcUserService) GetUser(c context.Context, req *inf.UserReq) (*inf.UserResp, error) {
	var user model.User
	err := p.DB.Gorm.Table(`users u`).Where(`u.id = ?`, req.Id).Scan(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("id:%d,get user err", req.Id))
	}
	return &inf.UserResp{Name: user.Name, Time: dongo_utils.Tick64()}, nil
}

func (p *RpcUserService) PushUser(context.Context, *inf.UserReq) (*empty.Empty, error) {
	return nil, nil
}
