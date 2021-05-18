package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Code2SessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCode2SessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) Code2SessionLogic {
	return Code2SessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Code2SessionLogic) Code2Session() (*types.LoginAppUser, error) {
	// todo: add your logic here and delete this line

	return &types.LoginAppUser{}, nil
}
