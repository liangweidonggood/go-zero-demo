package logic

import (
	"context"
	"go-zero-demo/datacenter/common/shared"
	"go-zero-demo/datacenter/service/common/cmd/rpc/common"
	"go-zero-demo/datacenter/service/common/model"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SnsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSnsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SnsInfoLogic {
	return SnsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SnsInfoLogic) SnsInfo(req types.SnsReq) (appConfigresp *common.AppConfigResp, err error) {
	//检查 缓存中是否有数据
	err = l.svcCtx.Cache.Get(model.GetCacheAppConfigIdPtyidPrefix(req.Beid.Beid, req.Ptyid), appConfigresp)
	if err != nil && err == shared.ErrNotFound {
		//直接请求 并返回
		appConfigresp, err = l.svcCtx.CommonRpc.GetAppConfig(l.ctx, &common.AppConfigReq{
			Beid:  req.Beid.Beid,
			Ptyid: req.Ptyid,
		})
		if err != nil {
			return
		}
		//缓存数据
		err = l.svcCtx.Cache.Set(model.GetCacheAppConfigIdPtyidPrefix(req.Beid.Beid, req.Ptyid), appConfigresp)
	}
	return
}
