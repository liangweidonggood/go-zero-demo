package handler

import (
	"go-zero-demo/datacenter/common/shared"
	"net/http"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/logic/common"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func WxTicketHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SnsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWxTicketLogic(r.Context(), ctx)
		resp, err := l.WxTicket(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			shared.OkJson(w, resp)
		}
	}
}
