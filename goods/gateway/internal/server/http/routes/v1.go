package routes

import (
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/handler/goods/category"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/handler/goods/group"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/handler/goods/item"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/handler/goods/itemsByCategoty"
	"github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/middleware"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

var MeRouteGroup = httpHelper.Routes{
	UrlPrefix: "/v1",
	Routes: []*httpHelper.Route{
		{
			UrlPattern: "/groups",
			Method:     httpHelper.GET_METHOD,
			Action:     group.Handle,
		},
		{
			UrlPattern: "/group/category",
			Method:     httpHelper.GET_METHOD,
			Action:     category.Handle,
		},
		{
			UrlPattern: "/group/category/items",
			Method:     httpHelper.GET_METHOD,
			Action:     itemsByCategoty.Handle,
		},
		{
			UrlPattern: "group/category/item",
			Method:     httpHelper.GET_METHOD,
			Action:     item.Handle,
		},
	},
	Middlewares: []httpHelper.Middleware{
		middleware.UserAuthenticationMiddleware,
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}
