package service

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetGrpcClientConnection(addr string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()
	clientConn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())

	if ctx.Err() == context.DeadlineExceeded {
		if clientConn != nil {
			clientConn.Close()
		}
		return nil, status.Error(codes.Code(503), "Create client connection timeout exceeded to addr: "+addr)
	}

	return clientConn, err
}
