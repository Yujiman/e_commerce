package types

import (
	"database/sql/driver"
	"strings"

	"github.com/Yujiman/e_commerce/auth/user/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PhoneType struct {
	name string
}

func NewPhoneType(value string) (*PhoneType, error) {
	if value == "" {
		return nil, status.Error(codes.Code(500), "The phone cannot be empty.")
	}
	if !utils.IsValidPhone(value) {
		return nil, status.Error(codes.Code(500), "Incorrect phone.")
	}
	return &PhoneType{name: strings.ToLower(value)}, nil
}

func (p PhoneType) Name() string {
	return p.name
}

func (p PhoneType) IsEqualTo(another PhoneType) bool {
	return p.name == another.name
}

// Value implements the driver.Valuer interface
func (p PhoneType) Value() (driver.Value, error) {
	if p.name == "" {
		return nil, nil
	}
	return p.name, nil
}

// Scan implements the sql.Scanner interface
func (p *PhoneType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	phoneType, err := NewPhoneType(src.(string))
	if err != nil {
		return status.Error(codes.Code(500), "Phone type from db has incorrect value")
	}
	*p = *phoneType

	return nil
}
