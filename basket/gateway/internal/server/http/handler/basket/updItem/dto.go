package updItem

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	UserId       string
	BasketItemId string
	Quantity     int64 `json:"quantity,omitempty"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.UserId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.BasketItemId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.Quantity,
			validation.Required,
		),
	)
}
