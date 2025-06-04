package service

import (
	"context"
	"errors"
	sqlRepository "github.com/danilkompanites/tinder-clone/services/match/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/match/pkg/model"
)

type Service struct {
	repo *sqlRepository.SqlRepository
}

func NewService(repo *sqlRepository.SqlRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) MakeDecisionById(ctx context.Context, match model.Match) (isMatch bool, err error) {
	res, err := s.repo.InsertMatch(ctx, match)
	if err != nil {
		return false, err
	}

	if res.FromDecision == res.ToDecision == true {
		return true, nil
	}

	return false, nil
}

func (s *Service) GetAllDecisionsByUserId(ctx context.Context, userId string) ([]*model.Match, error) {
	res, err := s.repo.SelectAllDecisionsByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, sqlRepository.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return res, nil
}
