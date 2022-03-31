package server

import (
	"net"
	"strconv"

	"github.com/Yujiman/e_commerce/goods/category/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/category/internal/proto/category"
	"github.com/Yujiman/e_commerce/goods/category/internal/utils"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCategoryServiceServer
}

func InitServer() {
	// Getting port
	port, err := config.GetServerPort()
	if err != nil {
		utils.LogFatalf("failed to get server address error: %v", err)
		return
	}

	// Listener
	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		utils.LogFatalf("failed to listen %v", err)
		return
	}

	// gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, &Server{})

	// Listen...
	err = grpcServer.Serve(listener)
	if err != nil {
		utils.LogFatalf("failed to serve %v", err)
		return
	}
}
