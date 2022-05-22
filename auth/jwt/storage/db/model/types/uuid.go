package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"github.com/Yujiman/e_commerce/auth/jwt/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UuidType struct {
	json.Marshaler
	json.Unmarshaler
	string sql.NullString
}

func NewUuidType(value string, null bool) (*UuidType, error) {
	if null {
		return &UuidType{string: sql.NullString{String: "", Valid: false}}, nil
	}
	if err := utils.CheckUuid(value); err != nil {
		return nil, status.Error(codes.Code(500), "Uuid value is invalid.")
	}
	return &UuidType{string: sql.NullString{String: value, Valid: true}}, nil
}

func (u UuidType) String() string {
	return u.string.String
}

func (u UuidType) IsEqualTo(another UuidType) bool {
	return u.string.String == another.string.String
}

// Value implements the driver.Valuer interface
func (u UuidType) Value() (driver.Value, error) {
	if !u.string.Valid {
		return nil, nil
	}
	return u.string.String, nil
}

// Scan implements the sql.Scanner interface
func (u *UuidType) Scan(src interface{}) error {
	if src == nil {
		*u = UuidType{string: sql.NullString{
			String: "",
			Valid:  false,
		}}
		return nil
	}
	uuidType, err := NewUuidType(src.(string), false)
	if err != nil {
		return status.Error(codes.Code(500), "Uuid type from db has incorrect value")
	}
	*u = *uuidType

	return nil
}

func (u UuidType) MarshalJSON() ([]byte, error) {
	if u.string.Valid {
		return json.Marshal(u.string.String)
	}
	return json.Marshal(nil)
}

func (u *UuidType) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		u.string.Valid = true
		u.string.String = *s
	} else {
		u.string.Valid = false
	}
	return nil
}
