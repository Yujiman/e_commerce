package refresh

import (
	"net/http"

	pbAggUser "github.com/Yujiman/e_commerce/auth/gateway/proto/aggregatorUser"
	pbAuthCheck "github.com/Yujiman/e_commerce/auth/gateway/proto/authentication"
	pbAuth "github.com/Yujiman/e_commerce/auth/gateway/proto/authorize"
	helperError "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/validator"
	"github.com/Yujiman/e_commerce/auth/gateway/service/aggregatorUser"
	"github.com/Yujiman/e_commerce/auth/gateway/service/authentication"
	"github.com/Yujiman/e_commerce/auth/gateway/service/authorize"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{}

	// Filling DTO from request
	dto.RefreshToken = request.Header.Get("x-satrap-2")

	// Validate DTO
	if !helperValidator.Validate(dto, response) {
		return
	}

	//Handler of gRPC
	data, err := authorize.AuthByRefresh(&pbAuth.AuthByRefreshRequest{RefreshToken: dto.RefreshToken})
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	tokenData, err := authentication.CheckAccess(&pbAuthCheck.CheckRequest{AccessToken: data.AccessToken})
	if err != nil {
		return
	}

	userData, err := aggregatorUser.GetById(&pbAggUser.GetByIdRequest{UserId: tokenData.UserId})
	if err != nil {
		return
	}

	response.Header().Add("X-Satrap-1", data.AccessToken)
	response.Header().Add("X-Satrap-2", data.RefreshToken)

	// Response
	helperHttp.JsonResponse(response, map[string]interface{}{
		"id":    userData.Id,
		"phone": userData.Phone,
		"email": userData.Email,
		"login": userData.Login,
	})
}
