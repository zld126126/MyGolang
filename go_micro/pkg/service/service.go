package service

import (
	"context"
	"fmt"

	proto "go_micro/proto"

	"go_micro/pkg/util"
)

type AssistService struct{}

func (p *AssistService) GetUser(ctx context.Context, req *proto.UserId, rsp *proto.User) error {
	rsp.Name = fmt.Sprintf(`%d`, req.Id)
	rsp.Time = util.ParseTimeToInt64()
	return nil
}

func (p *AssistService) GetActivity(ctx context.Context, req *proto.Name, rsp *proto.Activity) error {
	rsp.Name = fmt.Sprintf(`%s`, req.Name)
	rsp.Tp = proto.Tp_Tp_NotStart
	return nil
}
