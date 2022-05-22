package server

import (
	"context"

	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/handler/add"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/handler/remove"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/handler/update"
	pb "github.com/Yujiman/e_commerce/auth/dispatcherDomain/proto/dispatcherDomain"
	"github.com/Yujiman/e_commerce/auth/dispatcherDomain/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.UUID, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	resp, err := add.Handle(&add.RequestDTO{
		Name: req.Name,
		Url:  req.Url,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UUID{Value: resp.DomainId}, nil
}
func (Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	err := update.Handle(&update.RequestDTO{
		DomainId: req.DomainId,
		Name:     req.Name,
		Url:      req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
func (Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.Empty, error) {
	p, _ := peer.FromContext(ctx)
	clientIp := p.Addr.String()
	if !utils.CheckIp(clientIp) {
		return nil, status.Error(codes.Code(409), "Client IP "+clientIp+" not allowed")
	}

	err := remove.Handle(&remove.RemoveDTO{DomainId: req.DomainId})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
