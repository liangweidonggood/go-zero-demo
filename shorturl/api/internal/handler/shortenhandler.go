package handler

import (
	"net/http"

	"go-zero-demo/shorturl/api/internal/logic"
	"go-zero-demo/shorturl/api/internal/svc"
	"go-zero-demo/shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ShortenHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), ctx)
		resp, err := l.Shorten(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
