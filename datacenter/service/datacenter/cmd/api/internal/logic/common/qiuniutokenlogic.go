package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type QiuniuTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQiuniuTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) QiuniuTokenLogic {
	return QiuniuTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QiuniuTokenLogic) QiuniuToken(req types.Beid) (*types.Token, error) {
	// todo: add your logic here and delete this line

	return &types.Token{}, nil
}
