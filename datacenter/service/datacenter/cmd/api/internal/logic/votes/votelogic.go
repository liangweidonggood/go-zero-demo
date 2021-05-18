package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) VoteLogic {
	return VoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VoteLogic) Vote(req types.VoteReq) (*types.VoteResp, error) {
	// todo: add your logic here and delete this line

	return &types.VoteResp{}, nil
}
