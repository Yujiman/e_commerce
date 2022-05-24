package register

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/auth/gateway/server/http/helper/validator"
	"github.com/Yujiman/e_commerce/auth/gateway/service/dispatcherUser"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DTO struct {
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (dto *DTO) Validate() error {
	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.Phone,
			validation.Required,
			validation.Length(11, 11),
		),
		validation.Field(
			&dto.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&dto.Login,
			validation.Required,
			validation.Length(3, 20),
		),
		validation.Field(
			&dto.Password,
			validation.Required,
			validation.Length(3, 20),
		),
	)
}

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
	data, err := dispatcherUser.CreateUser(dto.Phone, dto.Email, dto.Login, dto.Password)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, map[string]interface{}{
		"id": data,
	})
}
