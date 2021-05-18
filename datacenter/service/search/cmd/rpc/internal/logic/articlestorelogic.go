package logic

import (
	"context"

	"go-zero-demo/datacenter/service/search/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/search/cmd/rpc/search"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleStoreLogic {
	return &ArticleStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleStoreLogic) ArticleStore(in *search.ArticleReq) (*search.Response, error) {
	// todo: add your logic here and delete this line

	return &search.Response{}, nil
}
