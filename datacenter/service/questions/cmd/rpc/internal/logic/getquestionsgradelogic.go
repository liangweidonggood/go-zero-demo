package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetQuestionsGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionsGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionsGradeLogic {
	return &GetQuestionsGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取得分
func (l *GetQuestionsGradeLogic) GetQuestionsGrade(in *questions.GradeReq) (*questions.QuestionsAnswerResp, error) {
	// todo: add your logic here and delete this line

	return &questions.QuestionsAnswerResp{}, nil
}
