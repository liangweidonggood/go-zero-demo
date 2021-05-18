package logic

import (
	"context"

	"go-zero-demo/datacenter/service/common/cmd/rpc/common"
	"go-zero-demo/datacenter/service/common/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAppConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppConfigLogic {
	return &GetAppConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppConfigLogic) GetAppConfig(in *common.AppConfigReq) (*common.AppConfigResp, error) {
	// todo: add your logic here and delete this line

	return &common.AppConfigResp{}, nil
}
