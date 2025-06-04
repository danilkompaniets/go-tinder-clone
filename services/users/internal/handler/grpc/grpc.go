package grpc

import (
	"context"
	"errors"
	"github.com/danilkompanites/tinder-clone/gen"
	"github.com/danilkompanites/tinder-clone/services/users/internal/service"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (h *Handler) SelectUsersByPreferences(ctx context.Context, req *gen.SelectUsersByPreferencesRequest) (*gen.SelectUsersByPreferencesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := h.service.GetUsersByPreferences(ctx, model.FromProtoToGetUsersByPreferencesRequest(req))

	if err != nil {
		return nil, err
	}

	var protoRes gen.SelectUsersByPreferencesResponse

	for _, user := range res {
		protoRes.Users = append(protoRes.Users, &gen.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: *user.FirstName,
			Bio:       *user.Bio,
			Gender:    *user.Gender,
			BirthDate: timestamppb.New(user.BirthDate),
			City:      user.City,
			AvatarUrl: user.AvatarURL,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}

	return &protoRes, nil
}
