package logic

import (
	"context"

	"go-zero-demo/datacenter/service/user/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SnsLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSnsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SnsLoginLogic {
	return &SnsLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SnsLoginLogic) SnsLogin(in *user.AppConfigReq) (*user.AppUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.AppUserResp{}, nil
}
