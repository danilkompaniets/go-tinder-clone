package grpc

import (
	"context"
	"errors"
	"github.com/danilkompanites/tinder-clone/gen"
	"github.com/danilkompanites/tinder-clone/services/users/internal/service"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/model"
	"google.golang.org/grpc"
	"time"
)

type Handler struct {
	gen.UnimplementedUserServer
	service service.UserService
}

func NewGRPCHandler(service service.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterGRPCHandler(server *grpc.Server, handler *Handler) {
	gen.RegisterUserServer(server, handler)
}

func (h *Handler) CreateUserFromAuth(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user := model.FromProtoToUser(req.User)

	if user == nil {
		return nil, errors.New("invalid user")
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := h.service.CreateUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &gen.CreateUserResponse{}, nil
}
