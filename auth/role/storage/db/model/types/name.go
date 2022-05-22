package types

import (
	"database/sql/driver"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NameType struct {
	name string
}

func NewNameType(value string) (*NameType, error) {
	if value == "" {
		return nil, status.Error(codes.Code(500), "The role name cannot be empty.")
	}
	if len(value) < 3 {
		return nil, status.Error(codes.Code(500), "Incorrect role name, min length = 3.")
	}

	return &NameType{name: strings.ToLower(value)}, nil
}

func (n NameType) Name() string {
	return n.name
}

func (n NameType) IsEqualTo(another NameType) bool {
	return n.name == another.name
}

// Value implements the driver.Valuer interface
func (n NameType) Value() (driver.Value, error) {
	if n.name == "" {
		return nil, nil
	}
	return n.name, nil
}

// Scan implements the sql.Scanner interface
func (n *NameType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	nameType, err := NewNameType(src.(string))
	if err != nil {
		return status.Error(codes.Code(500), "Name type from db has incorrect value")
	}
	*n = *nameType

	return nil
}
