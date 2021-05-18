package logic

import (
	"context"

	"go-zero-demo/datacenter/service/search/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/search/cmd/rpc/search"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleSearchLogic {
	return &ArticleSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleSearchLogic) ArticleSearch(in *search.SearchReq) (*search.ArticleResp, error) {
	// todo: add your logic here and delete this line

	return &search.ArticleResp{}, nil
}
