package logic

import (
	"context"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/logx"
)

type EnrollLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnrollLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnrollLogic {
	return &EnrollLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 报名
func (l *EnrollLogic) Enroll(in *votes.EnrollReq) (*votes.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &votes.EnrollResp{}, nil
}
