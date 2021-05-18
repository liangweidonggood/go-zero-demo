package logic

import (
	"context"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WxTicketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxTicketLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxTicketLogic {
	return WxTicketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxTicketLogic) WxTicket(req types.SnsReq) (*types.WxShareResp, error) {
	// todo: add your logic here and delete this line

	return &types.WxShareResp{}, nil
}
