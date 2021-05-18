// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	common "go-zero-demo/datacenter/service/datacenter/cmd/api/internal/handler/common"
	questions "go-zero-demo/datacenter/service/datacenter/cmd/api/internal/handler/questions"
	search "go-zero-demo/datacenter/service/datacenter/cmd/api/internal/handler/search"
	user "go-zero-demo/datacenter/service/datacenter/cmd/api/internal/handler/user"
	votes "go-zero-demo/datacenter/service/datacenter/cmd/api/internal/handler/votes"
	"go-zero-demo/datacenter/service/datacenter/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/common/appinfo",
				Handler: common.AppInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/common/snsinfo",
				Handler: common.SnsInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/common/wx/ticket",
				Handler: common.WxTicketHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/common/qiuniu/token",
				Handler: common.QiuniuTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/ping",
				Handler: user.PingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/wx/login",
				Handler: user.WxloginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/wx/login",
				Handler: user.Code2SessionHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Usercheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/dc/info",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/votes/activity/info",
				Handler: votes.ActivityInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/votes/activity/view",
				Handler: votes.ActivityIcrViewHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/votes/enroll/info",
				Handler: votes.EnrollInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/votes/enroll/lists",
				Handler: votes.EnrollListsHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Usercheck},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/votes/vote",
					Handler: votes.VoteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/votes/enroll",
					Handler: votes.EnrollHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Admincheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/search/article",
					Handler: search.ArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/search/articel/init",
					Handler: search.ArticleInitHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/search/articel/store",
					Handler: search.ArticleStoreHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/questions/activities/info",
				Handler: questions.ActivitiesInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/questions/award/info",
				Handler: questions.AwardInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/questions/award/list",
				Handler: questions.AwardListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Usercheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/questions/lists",
					Handler: questions.ListsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/questions/change",
					Handler: questions.ChangeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/questions/grade",
					Handler: questions.GradeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/questions/lottery/turntable",
					Handler: questions.TurntableHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/questions/lottery/convert",
					Handler: questions.LotteryHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
