package logic

import (
	"context"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetActivityInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityInfoLogic {
	return &GetActivityInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取 活动信息
func (l *GetActivityInfoLogic) GetActivityInfo(in *votes.ActInfoReq) (*votes.ActInfoResp, error) {
	// todo: add your logic here and delete this line

	return &votes.ActInfoResp{}, nil
}
