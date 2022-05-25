package routes

import (
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/handler/basket/createOrder"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/handler/basket/getBasket"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/handler/basket/putBasket"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/handler/basket/updItem"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/middleware"

	"github.com/autokz/go-http-server-helper/httpHelper"
)

var BasketRouteGroup = httpHelper.Routes{
	UrlPrefix: "/v1/basket",
	Routes: []*httpHelper.Route{
		{
			UrlPattern: "",
			Method:     httpHelper.GET_METHOD,
			Action:     getBasket.Handle,
		},
		{
			UrlPattern: "/put",
			Method:     httpHelper.PUT_METHOD,
			Action:     putBasket.Handle,
		},
		{
			UrlPattern: "/item/update",
			Method:     httpHelper.POST_METHOD,
			Action:     updItem.Handle,
		},
		{
			UrlPattern: "/order",
			Method:     httpHelper.POST_METHOD,
			Action:     createOrder.Handle,
		},
	},
	Middlewares: []httpHelper.Middleware{
		middleware.UserAuthenticationMiddleware,
		middleware.JsonMiddleware,
		middleware.CORSMiddleware,
	},
}
