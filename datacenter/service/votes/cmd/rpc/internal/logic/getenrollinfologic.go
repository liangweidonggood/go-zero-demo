package logic

import (
	"context"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetEnrollInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEnrollInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnrollInfoLogic {
	return &GetEnrollInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取作品信息
func (l *GetEnrollInfoLogic) GetEnrollInfo(in *votes.EnrollInfoReq) (*votes.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &votes.EnrollResp{}, nil
}
