package types

import (
	"database/sql/driver"
	"strings"

	"github.com/Yujiman/e_commerce/auth/user/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailType struct {
	name string
}

func NewEmailType(value string) (*EmailType, error) {
	if value == "" {
		return nil, status.Error(codes.Code(500), "The email cannot be empty.")
	}
	if !utils.IsValidEmail(value) {
		return nil, status.Error(codes.Code(500), "Incorrect email.")
	}
	return &EmailType{name: strings.ToLower(value)}, nil
}

func (e EmailType) Name() string {
	return e.name
}

func (e EmailType) IsEqualTo(another EmailType) bool {
	return e.name == another.name
}

// Value implements the driver.Valuer interface
func (e EmailType) Value() (driver.Value, error) {
	if e.name == "" {
		return nil, nil
	}
	return e.name, nil
}

// Scan implements the sql.Scanner interface
func (e *EmailType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	emailType, err := NewEmailType(src.(string))
	if err != nil {
		return status.Error(codes.Code(500), "Email type from db has incorrect value")
	}
	*e = *emailType

	return nil
}
