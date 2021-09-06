package grpc

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/database"
	"dongo_game_server/src/model"
	"dongo_game_server/src/util"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Grpc_UserService struct {
	DB *database.DB
}

func (p *Grpc_UserService) GetUser(c context.Context, req *inf.UserReq) (*inf.UserResp, error) {
	var user model.User
	err := p.DB.Gorm.Table(`users u`).Where(`u.id = ?`, req.Id).Scan(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("id:%d,get user err", req.Id))
	}
	return &inf.UserResp{Name: user.Name, Time: util.Tick64()}, nil
}

func (p *Grpc_UserService) PushUser(context.Context, *inf.UserReq) (*empty.Empty, error) {
	return nil, nil
}
