package helper

import (
	"net/http"

	httpHelper "github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/helper/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Validate(validatable validation.Validatable, responseWriter http.ResponseWriter) bool {
	err := validatable.Validate()
	if err != nil {
		httpHelper.ErrorResponse(err, responseWriter, http.StatusUnprocessableEntity)
		return false
	}
	return true
}
