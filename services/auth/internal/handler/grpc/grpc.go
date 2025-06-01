package grpcHandler

import (
	"context"
	"github.com/danilkompanites/tinder-clone/gen"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/service"
	"github.com/danilkompanites/tinder-clone/services/auth/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Handler struct {
	service *service.Service
	lis     net.Listener
	server  *grpc.Server
	gen.UnimplementedAuthServer
}

func NewHandler(service *service.Service, lis net.Listener) *Handler {
	handler := &Handler{service: service}
	server := grpc.NewServer()
	gen.RegisterAuthServer(server, handler)

	return &Handler{
		service: service,
		lis:     lis,
		server:  server,
	}
}

func (h *Handler) Run() error {
	err := h.server.Serve(h.lis)
	return err
}

func (h *Handler) RefreshToken(ctx context.Context, req *gen.RefreshTokenRequest) (*gen.RefreshTokenResponse, error) {
	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token is empty")
	}

	res, err := h.service.RefreshToken(ctx, &model.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &gen.RefreshTokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
