package deliveryPointUpdate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DTO struct {
	UserId          string `json:"user_id"`
	DeliveryPointId string `json:"delivery_point_id"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.UserId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.DeliveryPointId,
			validation.Required,
			is.UUIDv4,
		),
	)
}
