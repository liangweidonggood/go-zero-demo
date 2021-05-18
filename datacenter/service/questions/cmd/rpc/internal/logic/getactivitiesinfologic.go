package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetActivitiesInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivitiesInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivitiesInfoLogic {
	return &GetActivitiesInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取 问答抽奖活动信息
func (l *GetActivitiesInfoLogic) GetActivitiesInfo(in *questions.ActivitiesReq) (*questions.ActInfoResp, error) {
	// todo: add your logic here and delete this line

	return &questions.ActInfoResp{}, nil
}
