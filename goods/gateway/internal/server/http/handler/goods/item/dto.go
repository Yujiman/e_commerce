package item

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	ItemId string `json:"item_id"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.ItemId,
			validation.Required,
			is.UUIDv4,
		),
	)
}
