package createOrder

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/validator"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/middleware"
	serviceBasket "github.com/Yujiman/e_commerce/basket/gatway/internal/service/dispatherBasket"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{UserId: request.Header.Get(middleware.OauthUserId)}

	// Validate DTO
	if !helperValidator.Validate(dto, response) {
		return
	}

	//Handler of gRPC
	resp, err := serviceBasket.NewClient().CreateOrder(dto.UserId)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, resp)
}
