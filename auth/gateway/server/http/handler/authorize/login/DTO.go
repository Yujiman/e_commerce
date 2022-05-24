package login

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

func (dto *DTO) Validate() error {
	return validation.ValidateStruct(dto,
		validation.Field(
			&dto.Username,
			validation.Required,
			validation.Length(3, 0),
		),
		validation.Field(
			&dto.Password,
			validation.Required,
			validation.Length(6, 0),
		),
		validation.Field(
			&dto.Domain,
			validation.Required,
			validation.Length(3, 0),
		),
	)
}
