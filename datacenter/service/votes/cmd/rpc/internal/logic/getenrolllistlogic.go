package logic

import (
	"context"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetEnrollListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEnrollListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnrollListLogic {
	return &GetEnrollListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取 作品列表
func (l *GetEnrollListLogic) GetEnrollList(in *votes.ActidReq) (*votes.EnrollListResp, error) {
	// todo: add your logic here and delete this line

	return &votes.EnrollListResp{}, nil
}
