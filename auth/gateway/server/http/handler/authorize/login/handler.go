package login

import (
	"log"
	"net/http"

	pbAggUser "github.com/Yujiman/e_commerce/auth/gateway/proto/aggregatorUser"
	pbAuthCheck "github.com/Yujiman/e_commerce/auth/gateway/proto/authentication"
	pbAuth "github.com/Yujiman/e_commerce/auth/gateway/proto/authorize"
	helperError "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/validator"
	serviceAggregatorUser "github.com/Yujiman/e_commerce/auth/gateway/service/aggregatorUser"
	serviceAuthentication "github.com/Yujiman/e_commerce/auth/gateway/service/authentication"
	serviceAuthorize "github.com/Yujiman/e_commerce/auth/gateway/service/authorize"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{}

	// Filling DTO from request
	if !helperHttp.BindRequest(response, request, dto) {
		return
	}

	// Validate DTO
	if !helperValidator.Validate(dto, response) {
		return
	}

	//Handler of gRPC
	data, err := serviceAuthorize.AuthByPasswordDomain(&pbAuth.AuthByPasswordDomainRequest{
		Username:  dto.Username,
		Password:  dto.Password,
		DomainUrl: dto.Domain,
	})
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	tokenData, err := serviceAuthentication.CheckAccess(&pbAuthCheck.CheckRequest{AccessToken: data.AccessToken})
	if err != nil {
		return
	}

	userData, err := serviceAggregatorUser.GetById(&pbAggUser.GetByIdRequest{UserId: tokenData.UserId})
	if err != nil {
		return
	}

	log.Println(data)
	response.Header().Add("AccessToken", data.AccessToken)
	response.Header().Add("RefreshToken", data.RefreshToken)

	// Response
	helperHttp.JsonResponse(response, map[string]interface{}{
		"id":    userData.Id,
		"phone": userData.Phone,
		"email": userData.Email,
		"login": userData.Login,
	})
}
