package server

import (
	"log"
	"net"
	"strconv"

	"github.com/Yujiman/e_commerce/auth/user/internal/config"
	"github.com/Yujiman/e_commerce/auth/user/internal/proto/oauthUser"

	"google.golang.org/grpc"
)

type Server struct {
	oauthUser.UnimplementedOAuthUserServer
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

	grpcServer := grpc.NewServer()
	oauthUser.RegisterOAuthUserServer(grpcServer, &Server{})
	log.Println("Init Server. Start serving")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
		return
	}
}
