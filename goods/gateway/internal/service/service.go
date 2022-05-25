package service

import (
	"google.golang.org/grpc"
)

func GetGrpcClientConnection(addr string) (*grpc.ClientConn, error) {
	clientConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return clientConn, err
}
