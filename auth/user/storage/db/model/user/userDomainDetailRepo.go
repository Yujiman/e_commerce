package user

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/user/storage/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetDomainDetail(ctx context.Context, userId, domainId string) (*DomainDetail, error) {
	query := `select *
				from users_domains ud
				where user_id = $1 and  domain_id = $2;`

	dbConn := db.GetDbConnection()
	row := dbConn.QueryRowContext(ctx, query, userId, domainId)

	var domainDetail DomainDetail
	errNotFound := row.Scan(&domainDetail.UserId, &domainDetail.DomainId, &domainDetail.RoleId)

	if errNotFound != nil {
		return nil, status.Error(codes.Code(409), "User's domain detail not found;")
	}
	return &domainDetail, nil
}
