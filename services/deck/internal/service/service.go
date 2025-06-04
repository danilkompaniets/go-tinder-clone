package service

import (
	"context"
	"github.com/danilkompanites/tinder-clone/gen"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/model"
	"slices"
)

type Service struct {
	usersClient gen.UserClient
	matchClient gen.MatchClient
}

func NewService(usersClient gen.UserClient, matchClient gen.MatchClient) *Service {

	return &Service{
		usersClient: usersClient,
		matchClient: matchClient,
	}
}

func (s *Service) GetDeck(ctx context.Context, req *model.GetUsersByPreferencesRequest) ([]*model.User, error) {
	if req.Limit < 1 {
		req.Limit = 1
	}

	allUsers, err := s.usersClient.SelectUsersByPreferences(ctx, model.FromUserPreferencesRequestToProto(req))

	if err != nil {
		return nil, err
	}

	alreadyMatched, err := s.matchClient.GetDecisionsUserId(ctx, &gen.GetDecisionsUserIdRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	resLen := len(allUsers.Users) - len(alreadyMatched.UserIds)

	res := make([]*model.User, resLen)
	for _, usr := range allUsers.Users {
		isInvalid := slices.Contains(alreadyMatched.UserIds, usr.Id)
		if isInvalid {
			continue
		}
		res = append(res, model.FromProtoToUser(usr))
	}

	return res, nil
}
