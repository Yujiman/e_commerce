package deliveryPoint

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DTO struct {
	DeliveryPointId string `json:"delivery_point_id"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.DeliveryPointId,
			validation.Required,
			is.UUIDv4,
		),
	)
}
