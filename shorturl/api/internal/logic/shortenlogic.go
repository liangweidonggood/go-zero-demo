package logic

import (
	"context"
	"go-zero-demo/shorturl/rpc/transform/transformer"

	"go-zero-demo/shorturl/api/internal/svc"
	"go-zero-demo/shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (*types.ShortenResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return &types.ShortenResp{}, err
	}

	return &types.ShortenResp{
		Shorten: resp.Shorten,
	}, nil
	// 手动代码结束
}
