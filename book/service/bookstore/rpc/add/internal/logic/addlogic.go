package logic

import (
	"context"
	"go-zero-demo/book/service/bookstore/rpc/model"

	"go-zero-demo/book/service/bookstore/rpc/add/add"
	"go-zero-demo/book/service/bookstore/rpc/add/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// 手动代码开始
	_, err := l.svcCtx.Model.Insert(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &add.AddResp{
		Ok: true,
	}, nil
	// 手动代码结束
}
