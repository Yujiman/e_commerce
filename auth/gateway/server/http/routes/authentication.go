package routes

import (
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/handler/authentication/check"
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/middleware"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

var AuthCheck = httpHelper.Route{
	UrlPattern: "/v1/oauth/check",
	Method:     httpHelper.GET_METHOD,
	Action:     check.Handle,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}
