package itemsByCategoty

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	Category string `json:"category_id"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.Category,
			validation.Required,
			is.UUIDv4,
		),
	)
}
