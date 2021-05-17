package handler

import (
	"net/http"

	"go-zero-demo/shorturl/api/internal/logic"
	"go-zero-demo/shorturl/api/internal/svc"
	"go-zero-demo/shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ExpandHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), ctx)
		resp, err := l.Expand(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
