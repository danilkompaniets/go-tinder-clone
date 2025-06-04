package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/auth/internal/utils"
	"github.com/danilkompanites/tinder-clone/services/auth/pkg/model"
	userServiceModel "github.com/danilkompanites/tinder-clone/services/users/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	repository *sql.Repository
	cfg        *config.Config
	publisher  *producer.Publisher
}

func NewService(repository *sql.Repository, cfg *config.Config, publisher *producer.Publisher) *Service {
	return &Service{
		repository: repository,
		cfg:        cfg,
		publisher:  publisher,
	}
}

func (s *Service) RegisterUser(ctx context.Context, user *model.RegisterUserRequest) (*model.RegisterResponse, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id, err := s.repository.InsertUser(ctx, &model.User{
		Email:    user.Email,
		Password: string(hashedPass),
	})
	if err != nil {
		return nil, err
	}

	userCreatedEvent, err := json.Marshal(
		userServiceModel.User{
			ID:        id,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			Bio:       user.Bio,
			Gender:    user.Gender,
			BirthDate: user.BirthDate,
			City:      user.City,
			AvatarURL: user.AvatarURL,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	)

	if err != nil {
		return nil, err
	}

	s.publisher.PublishEvent(kafka.Events.UserCreated, id, userCreatedEvent)

	return &model.RegisterResponse{
		Id:    id,
		Email: user.Email,
	}, nil
}

func (s *Service) LoginUser(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	dbUser, err := s.repository.SelectUserByEmail(ctx, req.Email)
	if err != nil || dbUser == nil {
		fmt.Println(err)
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	accessToken, err := utils.CreateToken(dbUser.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.CreateToken(dbUser.ID, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	err = s.repository.InsertRefreshToken(ctx, dbUser.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (*model.RefreshTokenResponse, error) {
	if req.RefreshToken == "" {
		return nil, errors.New("refresh token is empty")
	}

	err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	user, err := s.repository.SelectUserByRefreshToken(ctx, req.RefreshToken)
	if err != nil || user == nil {
		return nil, errors.New("refresh token not found")
	}

	_ = s.repository.DeleteRefreshToken(ctx, req.RefreshToken)

	newRefreshToken, err := utils.CreateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	_ = s.repository.InsertRefreshToken(ctx, user.ID, newRefreshToken)

	newAccessToken, err := utils.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}

	return &model.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		return errors.New("refresh token is empty")
	}

	return s.repository.DeleteRefreshToken(ctx, refreshToken)
}

func (s *Service) GetUserByRefreshToken(ctx context.Context, tokenString string) (*model.User, error) {
	email, err := utils.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	res, err := s.repository.SelectUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return res, nil
}
