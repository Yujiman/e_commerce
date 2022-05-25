package getBasket

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/validator"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/middleware"
	serviceBasket "github.com/Yujiman/e_commerce/basket/gatway/internal/service/basket"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{UserId: request.Header.Get(middleware.OauthUserId)}

	// Filling DTO from request
	if !helperHttp.BindRequest(response, request, dto) {
		return
	}

	// Validate DTO
	if !helperValidator.Validate(dto, response) {
		return
	}

	//Handler of gRPC
	exist, err := serviceBasket.NewClient().ExistBasket(dto.UserId)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}
	if !exist {
		err := serviceBasket.NewClient().CreateBasket(dto.UserId)
		if err != nil {
			helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
			return
		}
	}

	basket, err := serviceBasket.NewClient().GetBasket(dto.UserId)
	if err != nil {
		return
	}

	// Response
	helperHttp.JsonResponse(response, basket)
}
