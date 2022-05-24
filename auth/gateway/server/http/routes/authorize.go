package routes

import (
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/handler/authorize/login"
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/handler/authorize/refresh"
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/handler/authorize/register"
	"github.com/Yujiman/e_commerce/auth/gateway/server/http/middleware"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

var AuthByRefresh = httpHelper.Route{
	UrlPattern: "/v1/oauth/refresh",
	Method:     httpHelper.POST_METHOD,
	Action:     refresh.Handle,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}

var AuthByPasswordDomain = httpHelper.Route{
	UrlPattern: "/v1/oauth/login",
	Method:     httpHelper.POST_METHOD,
	Action:     login.Handle,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}

var AuthRegister = httpHelper.Route{
	UrlPattern: "/v1/oauth/register",
	Method:     httpHelper.POST_METHOD,
	Action:     register.Handle,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}
