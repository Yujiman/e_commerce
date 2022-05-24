package routes

import (
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/cities"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/city"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/cityUpdate"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/deliveryPointUpdate"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/deliveryPoints"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/handler/me/me"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/middleware"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

var MeRouteGroup = httpHelper.Routes{
	UrlPrefix: "/v1",
	Routes: []*httpHelper.Route{
		{
			UrlPattern: "/delivery_points",
			Method:     httpHelper.GET_METHOD,
			Action:     deliveryPoints.Handle,
		},
		{
			UrlPattern: "/me/delivery_point/update",
			Method:     httpHelper.POST_METHOD,
			Action:     deliveryPointUpdate.Handle,
		},
		{
			UrlPattern: "/me/delivery_point",
			Method:     httpHelper.GET_METHOD,
			Action:     deliveryPoint.Handle,
		},
		{
			UrlPattern: "/me/city/update",
			Method:     httpHelper.POST_METHOD,
			Action:     cityUpdate.Handle,
		},
		{
			UrlPattern: "/city",
			Method:     httpHelper.GET_METHOD,
			Action:     city.Handle,
		},
		{
			UrlPattern: "/me",
			Method:     httpHelper.GET_METHOD,
			Action:     me.Handle,
		},
		{
			UrlPattern: "/cities",
			Method:     httpHelper.GET_METHOD,
			Action:     cities.Handle,
		},
	},
	Middlewares: []httpHelper.Middleware{
		middleware.UserAuthenticationMiddleware,
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}
