package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-demo/datacenter/service/user/cmd/api/internal/logic"
	"go-zero-demo/datacenter/service/user/cmd/api/internal/svc"
)

func userInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
