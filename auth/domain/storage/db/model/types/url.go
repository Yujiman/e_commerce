package types

import (
	"database/sql/driver"
	"strings"

	"github.com/Yujiman/e_commerce/auth/domain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UrlType struct {
	url string
}

func NewUrlType(value string) (*UrlType, error) {
	if value == "" {
		return nil, status.Error(codes.Code(500), "The url cannot be empty.")
	}
	if !utils.IsValidUrl(value) {
		return nil, status.Error(codes.Code(500), "Incorrect url.")
	}

	return &UrlType{url: strings.ToLower(value)}, nil
}

func (u UrlType) Url() string {
	return u.url
}

func (u UrlType) IsEqualTo(another UrlType) bool {
	return u.url == another.url
}

// Value implements the driver.Valuer interface
func (u UrlType) Value() (driver.Value, error) {
	if u.url == "" {
		return nil, nil
	}
	return u.url, nil
}

// Scan implements the sql.Scanner interface
func (u *UrlType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	urlType, err := NewUrlType(src.(string))
	if err != nil {
		return status.Error(codes.Code(500), "Url type from db has incorrect value")
	}
	*u = *urlType

	return nil
}
