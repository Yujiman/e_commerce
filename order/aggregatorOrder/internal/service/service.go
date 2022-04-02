package service

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetGrpcClientConnection(addr string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientConn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())

	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Code(503), "Create client connection timeout exceeded by addr: "+addr)
	}

	return clientConn, err
}
