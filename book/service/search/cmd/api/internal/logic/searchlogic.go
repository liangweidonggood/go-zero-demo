package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-demo/book/service/user/cmd/rpc/userclient"

	"go-zero-demo/book/service/search/cmd/api/internal/svc"
	"go-zero-demo/book/service/search/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
	//解析出来jwt里的userid
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 使用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	//todo 调用本地服务返回结果

	return &types.SearchReply{
		Name:  req.Name,
		Count: 100,
	}, nil
}
