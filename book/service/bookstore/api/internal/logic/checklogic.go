package logic

import (
	"context"
	"go-zero-demo/book/service/bookstore/rpc/check/checker"

	"go-zero-demo/book/service/bookstore/api/internal/svc"
	"go-zero-demo/book/service/bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: req.Book,
	})
	if err != nil {
		logx.Error(err)
		return &types.CheckResp{}, err
	}

	return &types.CheckResp{
		Found: resp.Found,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
