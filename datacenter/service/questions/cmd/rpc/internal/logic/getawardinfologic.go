package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAwardInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAwardInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAwardInfoLogic {
	return &GetAwardInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取 问答奖品信息
func (l *GetAwardInfoLogic) GetAwardInfo(in *questions.ActivitiesReq) (*questions.AwardInfoResp, error) {
	// todo: add your logic here and delete this line

	return &questions.AwardInfoResp{}, nil
}
