package logic

import (
	"context"

	"go-zero-demo/datacenter/service/questions/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questions"

	"github.com/tal-tech/go-zero/core/logx"
)

type PostConvertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostConvertLogic {
	return &PostConvertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 填写中奖记录
func (l *PostConvertLogic) PostConvert(in *questions.ConvertReq) (*questions.ConvertResp, error) {
	// todo: add your logic here and delete this line

	return &questions.ConvertResp{}, nil
}
