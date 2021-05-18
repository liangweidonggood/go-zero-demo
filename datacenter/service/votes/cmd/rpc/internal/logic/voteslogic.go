package logic

import (
	"context"

	"go-zero-demo/datacenter/service/votes/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/votes/cmd/rpc/votes"

	"github.com/tal-tech/go-zero/core/logx"
)

type VotesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVotesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VotesLogic {
	return &VotesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 投票
func (l *VotesLogic) Votes(in *votes.VotesReq) (*votes.VotesResp, error) {
	// todo: add your logic here and delete this line

	return &votes.VotesResp{}, nil
}
