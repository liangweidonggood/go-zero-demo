package handler

import (
	"go-zero-demo/datacenter/common/shared"
	"go-zero-demo/datacenter/service/questions/cmd/rpc/questionsclient"
	"net/http"

	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/logic/questions"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ActivitiesInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		actid := shared.StrToInt64(query.Get("actid"))
		activitReq := questionsclient.ActivitiesReq{
			Actid: actid,
		}

		l := logic.NewActivitiesInfoLogic(r.Context(), ctx)
		resp, err := l.ActivitiesInfo(activitReq)
		if err != nil {
			httpx.Error(w, err)
		} else {
			shared.OkJson(w, resp)
		}
	}
}
