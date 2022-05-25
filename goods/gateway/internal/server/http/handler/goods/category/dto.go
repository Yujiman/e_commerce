package category

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	GroupId string `json:"group_id"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.GroupId,
			validation.Required,
			is.UUIDv4,
		),
	)
}
