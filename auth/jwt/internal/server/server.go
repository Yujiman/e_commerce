package server

import (
	"log"
	"net"
	"strconv"

	"github.com/Yujiman/e_commerce/auth/jwt/internal/config"
	"github.com/Yujiman/e_commerce/auth/jwt/internal/proto/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Server struct {
	jwt.UnimplementedJwtServer
	Name string
	Keys *config.Keys
}

func getServerAddress() (string, error) {

	port, err := config.GetServerPort()
	if err != nil {
		return "", err
	}

	addr := "0.0.0.0:" + strconv.Itoa(port)

	return addr, nil
}

func InitServer() {
	addr, err := getServerAddress()
	if err != nil {
		log.Fatalf("failed to get server address error: %v", err)
		return
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
		return
	}

	keysStorage := config.GetKeysStorage()
	keys := &config.Keys{Storage: keysStorage}

	server := &Server{
		Keys: keys,
		Name: "jwt-service",
	}
	grpcServer := grpc.NewServer()
	jwt.RegisterJwtServer(grpcServer, server)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
		return
	}
}

func (s Server) handleServerError(err error) error {
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.Unknown || statusErr.Code() == 500 {
				return status.Error(statusErr.Code(), s.Name+": "+statusErr.Message())
			}
		}
		return err
	}
	return nil
}
