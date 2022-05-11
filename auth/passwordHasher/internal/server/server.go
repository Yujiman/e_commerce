package server

import (
	context "context"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/Yujiman/e_commerce/auth/passwordHasher/internal/config"
	"github.com/Yujiman/e_commerce/auth/passwordHasher/internal/handler"
	"github.com/Yujiman/e_commerce/auth/passwordHasher/internal/proto/passwordHasher"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type Server struct {
	passwordHasher.UnimplementedPasswordHashServer
	params *config.Argon2Params
}

func (s *Server) CreateHash(ctx context.Context, request *passwordHasher.CreateHashRequest) (*passwordHasher.CreateHashResponse, error) {
	defer ctx.Done()
	password := strings.Trim(request.Password, " ")
	if password == "" {
		return nil, status.Error(codes.Code(400), "Password string couldn't be an empty.")
	}
	hash, err := handler.CreateHash(password, s.params)

	if err != nil {
		return nil, err
	}

	return &passwordHasher.CreateHashResponse{Hash: hash}, nil
}

func (s *Server) Validate(ctx context.Context, request *passwordHasher.ValidateRequest) (*passwordHasher.ValidateResponse, error) {
	defer ctx.Done()
	password := strings.Trim(request.Password, " ")
	if password == "" {
		return nil, status.Error(codes.Code(400), "Password string couldn't be an empty.")
	}
	hash := strings.Trim(request.Hash, " ")
	if password == "" {
		return nil, status.Error(codes.Code(400), "Password hash couldn't be an empty.")
	}
	valid, err := handler.Validate(password, hash)
	if err != nil {
		return nil, err
	}

	return &passwordHasher.ValidateResponse{Valid: valid}, nil
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

	params := config.GetDefaultArgon2idParams()

	grpcServer := grpc.NewServer()
	passwordHasher.RegisterPasswordHashServer(grpcServer, &Server{params: params})

	log.Println("Started")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
		return
	}

}
