package createOrder

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	UserId string
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.UserId,
			validation.Required,
			is.UUIDv4,
		),
	)
}
