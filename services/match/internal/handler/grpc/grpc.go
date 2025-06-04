package grpcHandler

import (
	"context"
	"errors"
	"github.com/danilkompanites/tinder-clone/gen"
	sqlRepository "github.com/danilkompanites/tinder-clone/services/match/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/match/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	service *service.Service
	*gen.UnimplementedMatchServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMyLikesByUserId(ctx context.Context, req *gen.GetDecisionsUserIdRequest) (*gen.GetDecisionsUserIdResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}

	decisions, err := h.service.GetAllDecisionsByUserId(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, sqlRepository.ErrNotFound) {
			return nil, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]string, len(decisions))

	for _, decision := range decisions {
		res = append(res, decision.ToId)
	}

	return &gen.GetDecisionsUserIdResponse{UserIds: res}, nil
}
