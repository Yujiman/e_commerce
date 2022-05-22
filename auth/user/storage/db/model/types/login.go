package types

import (
	"database/sql/driver"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginType struct {
	name string
}

func NewLoginType(value string) (*LoginType, error) {
	if value == "" {
		return nil, status.Error(codes.Code(500), "The login value cannot be empty.")
	}
	if len(value) < 3 {
		return nil, status.Error(codes.Code(500), "Incorrect login value, length min=3.")
	}
	return &LoginType{name: strings.ToLower(value)}, nil
}

func (l LoginType) Name() string {
	return l.name
}

func (l LoginType) IsEqualTo(another LoginType) bool {
	return l.name == another.name
}

// Value implements the driver.Valuer interface
func (l LoginType) Value() (driver.Value, error) {
	if l.name == "" {
		return nil, nil
	}
	return l.name, nil
}

// Scan implements the sql.Scanner interface
func (l *LoginType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	loginType, err := NewLoginType(src.(string))
	if err != nil {
		return status.Error(codes.Code(500), "Login type from db has incorrect value")
	}
	*l = *loginType

	return nil
}
