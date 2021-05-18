package logic

import (
	"context"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/logx"
)

type IncrActiviViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrActiviViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrActiviViewLogic {
	return &IncrActiviViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加活动的爆光量
func (l *IncrActiviViewLogic) IncrActiviView(in *votes.ActInfoReq) (*votes.ActInfoResp, error) {
	// todo: add your logic here and delete this line

	return &votes.ActInfoResp{}, nil
}
