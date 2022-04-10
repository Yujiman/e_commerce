package cityUpdate

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	CityId string `json:"city_id"`
	UserId string `json:"user_id"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.CityId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.UserId,
			validation.Required,
			is.UUIDv4,
		),
	)
}
