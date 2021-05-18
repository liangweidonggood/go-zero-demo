package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type PostTurnTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostTurnTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostTurnTableLogic {
	return &PostTurnTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 抽奖
func (l *PostTurnTableLogic) PostTurnTable(in *questions.TurnTableReq) (*questions.AwardInfoResp, error) {
	// todo: add your logic here and delete this line

	return &questions.AwardInfoResp{}, nil
}
