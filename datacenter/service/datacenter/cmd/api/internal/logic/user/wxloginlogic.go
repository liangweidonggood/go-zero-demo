package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WxloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxloginLogic {
	return WxloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxloginLogic) Wxlogin(req types.WxLoginReq) (*types.LoginAppUser, error) {
	// todo: add your logic here and delete this line

	return &types.LoginAppUser{}, nil
}
