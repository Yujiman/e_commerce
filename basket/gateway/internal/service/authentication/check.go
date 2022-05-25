package authentication

import (
	"context"
	"time"

	"github.com/Yujiman/e_commerce/basket/gatway/internal/config"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/service"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/utils"

	pbAuthentication "github.com/Yujiman/e_commerce/basket/gatway/internal/proto/authentication"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckAccess(request *pbAuthentication.CheckRequest) (*pbAuthentication.TokenData, error) {
	// Create gRPC connection
	addr := config.GetConfig().ServicesParam.Authentication

	connection, err := service.GetGrpcClientConnection(addr)
	defer utils.MuteCloseClientConn(connection)
	if err != nil {
		return nil, err
	}

	// Create timeout for call gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()

	// Call gRPC method...
	client := pbAuthentication.NewAuthenticationServiceClient(connection)
	resp, err := client.Check(ctx, request)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Client to Gateway->Authentication:CheckAccess service timeout exceeded.")
	}

	if err != nil {
		return nil, err
	}

	return resp, err
}
