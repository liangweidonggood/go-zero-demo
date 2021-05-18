package logic

import (
	"context"

	"go-zero-demo/datacenter/service/user/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.UserReply, error) {
	// todo: add your logic here and delete this line

	return &user.UserReply{}, nil
}
