package logic

import (
	"context"

	"go-zero-demo/book/service/user/cmd/api/internal/svc"
	"go-zero-demo/book/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
	// todo: add your logic here and delete this line

	return &types.LoginReply{}, nil
}
