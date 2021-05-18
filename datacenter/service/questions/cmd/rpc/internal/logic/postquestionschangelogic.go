package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type PostQuestionsChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostQuestionsChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostQuestionsChangeLogic {
	return &PostQuestionsChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  提交 答案
func (l *PostQuestionsChangeLogic) PostQuestionsChange(in *questions.QuestionsAnswerReq) (*questions.QuestionsAnswerResp, error) {
	// todo: add your logic here and delete this line

	return &questions.QuestionsAnswerResp{}, nil
}
