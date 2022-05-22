package refresh

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DTO struct {
	RefreshToken string `json:"x-satrap-2"`
}

func (dto *DTO) Validate() error {
	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.RefreshToken,
			validation.Required,
			validation.Length(10, 0),
		),
	)
}
