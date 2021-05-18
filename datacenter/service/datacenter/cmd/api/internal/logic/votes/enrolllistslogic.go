package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EnrollListsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnrollListsLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnrollListsLogic {
	return EnrollListsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnrollListsLogic) EnrollLists(req types.Actid) (*types.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &types.EnrollResp{}, nil
}
