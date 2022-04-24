package types

import (
	"database/sql/driver"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const StatusWait = "wait"
const StatusActive = "active"
const StatusLocked = "locked"

type StatusType struct {
	name string
}

func NewStatus(value string) (*StatusType, error) {
	if value != StatusWait && value != StatusActive && value != StatusLocked {
		return nil, status.Error(codes.Code(500), "Status value invalid.")
	}
	return &StatusType{name: value}, nil
}

func (s StatusType) EqualsTo(another StatusType) bool {
	return strings.EqualFold(s.name, another.name)
}

func (s StatusType) Wait() *StatusType {
	s.name = StatusWait
	return &s
}

func (s StatusType) Active() *StatusType {
	s.name = StatusActive
	return &s
}

func (s StatusType) Locked() *StatusType {
	s.name = StatusLocked
	return &s
}

func (s StatusType) IsWait() bool {
	return s.name == StatusWait
}

func (s StatusType) IsActive() bool {
	return s.name == StatusActive
}

func (s StatusType) IsLocked() bool {
	return s.name == StatusLocked
}

func (s StatusType) String() string {
	return s.name
}

// Value implements the driver.Valuer interface
func (s StatusType) Value() (driver.Value, error) {
	if s.String() == "" {
		return nil, status.Error(codes.Code(500), "Status type couldn't be empty.")
	}
	return s.String(), nil
}

// Scan implements the sql.Scanner interface
func (s *StatusType) Scan(src interface{}) error {
	statusType, err := NewStatus(src.(string))
	if err != nil {
		return status.Error(codes.Code(500), "Status type from db has incorrect value")
	}

	*s = *statusType
	return nil
}
