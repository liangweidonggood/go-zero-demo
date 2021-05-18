package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EnrollInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnrollInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnrollInfoLogic {
	return EnrollInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnrollInfoLogic) EnrollInfo(req types.VoteReq) (*types.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &types.EnrollResp{}, nil
}
