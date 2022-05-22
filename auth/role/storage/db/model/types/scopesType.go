package types

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"github.com/Yujiman/e_commerce/auth/role/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScopesType struct {
	string sql.NullString
}

func NewScopesType(value string, null bool) (*ScopesType, error) {
	if null {
		return &ScopesType{string: sql.NullString{String: "", Valid: false}}, nil
	}
	if !utils.IsValidJson(value) {
		return nil, status.Error(codes.Code(500), "Scopes value is invalid, must be json string.")
	}
	buf := new(bytes.Buffer)
	_ = json.Compact(buf, []byte(value))
	value = buf.String()
	return &ScopesType{string: sql.NullString{String: value, Valid: true}}, nil
}

func (s ScopesType) String() string {
	return s.string.String
}

func (s ScopesType) IsEqualTo(another ScopesType) bool {
	return s.string.String == another.string.String
}

// Value implements the driver.Valuer interface
func (s ScopesType) Value() (driver.Value, error) {
	if !s.string.Valid {
		return nil, nil
	}
	return s.string.String, nil
}

// Scan implements the sql.Scanner interface
func (s *ScopesType) Scan(src interface{}) error {
	if src == nil {
		*s = ScopesType{string: sql.NullString{
			String: "",
			Valid:  false,
		}}
		return nil
	}
	scopesType, err := NewScopesType(src.(string), false)
	if err != nil {
		return status.Error(codes.Code(500), "Scopes type from db has incorrect value")
	}
	*s = *scopesType

	return nil
}
