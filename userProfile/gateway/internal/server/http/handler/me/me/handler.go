package me

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/validator"
	serviceuser "github.com/Yujiman/e_commerce/userProfile/gatway/internal/service/user"
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
	data, err := serviceuser.GetUserInfo(dto.UserId)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, data)
}
