package check

import (
	"net/http"

	pbAggUser "github.com/Yujiman/e_commerce/auth/gateway/proto/aggregatorUser"
	pbAuthentication "github.com/Yujiman/e_commerce/auth/gateway/proto/authentication"
	helperError "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/validator"
	serviceAggregatorUser "github.com/Yujiman/e_commerce/auth/gateway/service/aggregatorUser"
	serviceAuthentication "github.com/Yujiman/e_commerce/auth/gateway/service/authentication"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{}

	// Filling DTO from request
	dto.AccessToken = request.Header.Get("x-satrap-1")

	// Validate DTO
	if !helperValidator.Validate(dto, response) {
		return
	}

	//Handler of gRPC
	tokenData, err := serviceAuthentication.CheckAccess(&pbAuthentication.CheckRequest{AccessToken: dto.AccessToken})
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	userData, err := serviceAggregatorUser.GetById(&pbAggUser.GetByIdRequest{UserId: tokenData.UserId})
	if err != nil {
		return
	}

	// Response
	helperHttp.JsonResponse(response, map[string]interface{}{
		"id":    userData.Id,
		"phone": userData.Phone,
		"email": userData.Email,
		"login": userData.Login,
	})
}
