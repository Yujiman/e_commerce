package putBasket

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	UserId   string
	BasketId string
	GoodId   string  `json:"good_id,omitempty"`
	Quantity int64   `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

func (dto *DTO) Validate() error {

	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.UserId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.BasketId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.GoodId,
			validation.Required,
			is.UUIDv4,
		),
		validation.Field(
			&dto.Quantity,
			validation.Required,
		),
	)
}
