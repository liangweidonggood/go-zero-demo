package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAwardListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAwardListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAwardListLogic {
	return &GetAwardListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取 问答奖品列表
func (l *GetAwardListLogic) GetAwardList(in *questions.ActivitiesReq) (*questions.AwardListResp, error) {
	// todo: add your logic here and delete this line

	return &questions.AwardListResp{}, nil
}
