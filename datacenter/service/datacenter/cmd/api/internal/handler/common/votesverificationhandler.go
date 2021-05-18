package handler

import (
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"
	"net/http"
)

func VotesVerificationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("NT04cqknJe0em3mT"))
		return
	}
}
