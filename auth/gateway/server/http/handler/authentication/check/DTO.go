package check

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DTO struct {
	AccessToken string `json:"x-satrap-1"`
}

func (dto *DTO) Validate() error {
	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.AccessToken,
			validation.Required,
			validation.Length(10, 0),
		),
	)
}
