package logic

import (
	"context"

	"go-zero-demo/datacenter/service/common/cmd/rpc/common"
	"go-zero-demo/datacenter/service/common/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBaseAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBaseAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseAppLogic {
	return &GetBaseAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBaseAppLogic) GetBaseApp(in *common.BaseAppReq) (*common.BaseAppResp, error) {
	repl, err := l.svcCtx.BaseAppModel.FindOne(in.Beid)
	if err != nil {
		return nil, err
	}

	return &common.BaseAppResp{
		Beid:        repl.Id,
		Logo:        repl.Logo,
		Fullwebsite: repl.Fullwebsite,
		Isclose:     repl.Isclose,
	}, nil
}
