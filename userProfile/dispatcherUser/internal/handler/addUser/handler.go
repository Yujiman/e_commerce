package addUser

import (
	"context"

	pb "github.com/Yujiman/e_commerce/userProfile/dispatcherUser/internal/proto/dispatcherUser"
)

func Handle(ctx context.Context, request *pb.AddUserRequest) (*pb.UUID, error) {
	// TODO Implement your handler logic!
	return nil, nil
}
