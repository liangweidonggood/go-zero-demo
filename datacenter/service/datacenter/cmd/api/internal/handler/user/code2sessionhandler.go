package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/logic/user"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
)

func Code2SessionHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewCode2SessionLogic(r.Context(), ctx)
		resp, err := l.Code2Session()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
