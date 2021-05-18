package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetQuestionsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionsListLogic {
	return &GetQuestionsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取 问题列表
func (l *GetQuestionsListLogic) GetQuestionsList(in *questions.ActivitiesReq) (*questions.QuestionsListResp, error) {
	// todo: add your logic here and delete this line

	return &questions.QuestionsListResp{}, nil
}
