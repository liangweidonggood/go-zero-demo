package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListsLogic {
	return ListsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListsLogic) Lists(req types.VoteReq) (*types.AnswerResp, error) {
	// todo: add your logic here and delete this line

	return &types.AnswerResp{}, nil
}
