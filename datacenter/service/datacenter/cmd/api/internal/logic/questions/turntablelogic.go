package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TurntableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTurntableLogic(ctx context.Context, svcCtx *svc.ServiceContext) TurntableLogic {
	return TurntableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TurntableLogic) Turntable(req types.EnrollReq) (*types.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &types.EnrollResp{}, nil
}
