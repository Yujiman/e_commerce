package types

import (
	"database/sql"
	"database/sql/driver"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	New        = "new"
	InProgress = "in_progress"
	InPoint    = "in_point"
	Issued     = "issued"
)

type StatusType struct {
	val sql.NullString
}

func NewStatusType(value string, null bool) (*StatusType, error) {
	if null {
		return &StatusType{val: sql.NullString{String: "", Valid: false}}, nil
	}

	value = strings.ToLower(value)
	if !IsStatusType(value) {
		return nil, status.Error(codes.Code(409), "Invalid status code.")
	}

	return &StatusType{val: sql.NullString{String: value, Valid: true}}, nil
}

func IsStatusType(value string) bool {
	value = strings.ToLower(value)
	types := []string{New, InProgress, InPoint, Issued}

	return stringInSlice(value, types)
}

func (s StatusType) String() string {
	if !s.val.Valid {
		return ""
	}
	return s.val.String
}

func (s StatusType) IsNew() bool {
	return s.val.String == New
}

func (s StatusType) IsEqualTo(another StatusType) bool {
	return s.val.String == another.val.String
}

// Value implements the driver.Valuer interface
func (s StatusType) Value() (driver.Value, error) {
	if !s.val.Valid {
		return nil, nil
	}
	return s.val.String, nil
}

// Scan implements the sql.Scanner interface
func (s *StatusType) Scan(src interface{}) error {
	if src == nil {
		*s = StatusType{val: sql.NullString{
			String: "",
			Valid:  false,
		}}
		return nil
	}

	sFromDB, err := NewStatusType(src.(string), false)
	if err != nil {
		return status.Error(codes.Code(500), "Status type from db has incorrect value.")
	}
	*s = *sFromDB

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
