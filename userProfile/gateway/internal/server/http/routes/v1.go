package routes

import (
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/city"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/cityUpdate"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/deliveryPointUpdate"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/me"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/middleware"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

var MeRouteGroup = httpHelper.Routes{
	UrlPrefix: "/v1/me",
	Routes: []*httpHelper.Route{
		{
			UrlPattern: "/delivery_point/update",
			Method:     httpHelper.POST_METHOD,
			Action:     deliveryPointUpdate.Handle,
		},
		{
			UrlPattern: "/delivery_point",
			Method:     httpHelper.GET_METHOD,
			Action:     deliveryPoint.Handle,
		},
		{
			UrlPattern: "/city/update",
			Method:     httpHelper.POST_METHOD,
			Action:     cityUpdate.Handle,
		},
		{
			UrlPattern: "/city",
			Method:     httpHelper.GET_METHOD,
			Action:     city.Handle,
		},
		{
			UrlPattern: "",
			Method:     httpHelper.GET_METHOD,
			Action:     me.Handle,
		},
	},
	Middlewares: []httpHelper.Middleware{
		//middleware.UserAuthenticationMiddleware,
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}
