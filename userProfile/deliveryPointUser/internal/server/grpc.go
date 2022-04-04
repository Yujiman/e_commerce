package server

import (
	"net"
	"strconv"

	"github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/config"
	pb "github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/proto/deliveryPointUser"
	"github.com/Yujiman/e_commerce/goods/userProfile/deliveryPointUser/internal/utils"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDeliveryPointUserServiceServer
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
	pb.RegisterDeliveryPointUserServiceServer(grpcServer, &Server{})

	// Listen...
	err = grpcServer.Serve(listener)
	if err != nil {
		utils.LogFatalf("failed to serve %v", err)
		return
	}
}
