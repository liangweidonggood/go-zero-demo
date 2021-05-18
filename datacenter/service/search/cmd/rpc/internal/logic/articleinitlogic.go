package logic

import (
	"context"

	"go-zero-demo/datacenter/service/search/cmd/rpc/internal/svc"
	"go-zero-demo/datacenter/service/search/cmd/rpc/search"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleInitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleInitLogic {
	return &ArticleInitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleInitLogic) ArticleInit(in *search.Request) (*search.Response, error) {
	// todo: add your logic here and delete this line

	return &search.Response{}, nil
}
