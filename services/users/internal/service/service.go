package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/danilkompanites/tinder-clone/internal/config"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka"
	"github.com/danilkompanites/tinder-clone/internal/utils/kafka/producer"
	"github.com/danilkompanites/tinder-clone/internal/utils/storage/s3"
	"github.com/danilkompanites/tinder-clone/services/users/internal/repository/sql"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/model"
	"github.com/minio/minio-go/v7"
)

type UserService struct {
	cfg       *config.Config
	repo      *sql.Repository
	publisher *producer.Publisher
	s3Client  *s3.Client
}

func NewUserService(repo *sql.Repository, cfg *config.Config, publisher *producer.Publisher) *UserService {
	return &UserService{
		repo:      repo,
		cfg:       cfg,
		publisher: publisher,
	}
}

func (s *UserService) AddUserPhoto(ctx context.Context, userId string, req model.AddUserPhotoRequest) (string, error) {
	if req.Photo == nil {
		return "", errors.New("photo is required")
	}
	photo := req.Photo
	buffer, err := photo.Open()
	if err != nil {
		return "", err
	}

	info, err := s.s3Client.PutObject(ctx, s.cfg.Services.ObjectStorage.S3, photo.Filename, buffer, photo.Size, minio.PutObjectOptions{
		ContentType: photo.Header.Get("Content-Type"),
	})

	if err != nil {
		return "", err
	}

	err = s.repo.InsertUserPhotoURL(ctx, info.Location, userId, *req.Position)

	if err != nil {
		return "", err
	}

	return info.Location, nil
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) error {
	if user.ID == "" || user.Email == "" || user.Username == "" {
		return errors.New("missing required user fields")
	}
	return s.repo.InsertUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("user id is empty")
	}
	return s.repo.SelectUserByID(ctx, id)
}

func (s *UserService) CreateUserPreference(ctx context.Context, req model.CreatePreferenceRequest) error {
	if req.UserId == "" {
		return errors.New("user id is empty")
	}

	id, err := s.repo.InsertUserPreference(ctx, req)
	if err != nil {
		return err
	}

	prefs := kafka.PreferencesUpdate{
		ID:              id,
		PreferredGender: &req.PreferredGender,
		AgeMax:          &req.AgeMax,
		AgeMin:          &req.AgeMin,
		Position:        req.Position,
	}

	prefsJson, err := json.Marshal(prefs)
	if err != nil {
		return err
	}

	s.publisher.PublishEvent(kafka.Events.UserPreferencesUpdated, prefs.ID, prefsJson)

	return nil
}

func (s *UserService) UpdateUserPreference(ctx context.Context, req model.UserPreferences) error {
	if req.ID == "" {
		return errors.New("user id is empty")
	}

	prefs := kafka.PreferencesUpdate{
		ID:              req.ID,
		PreferredGender: req.PreferredGender,
		AgeMax:          req.AgeMax,
		AgeMin:          req.AgeMin,
		Position:        req.Position,
	}

	prefsJson, err := json.Marshal(prefs)
	if err != nil {
		return err
	}

	s.publisher.PublishEvent(kafka.Events.UserPreferencesUpdated, prefs.ID, prefsJson)

	return s.repo.UpdateUserPreference(ctx, req)
}

func (s *UserService) DeleteUserPreference(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("preference id is empty")
	}

	return s.repo.DeleteUserPreference(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user model.UpdateUserRequest) error {
	if user.ID == "" {
		return errors.New("missing required user fields")
	}

	return s.repo.UpdateUser(ctx, &user)
}

func (s *UserService) GetUsersByPreferences(ctx context.Context, req *model.GetUsersByPreferencesRequest) ([]*model.User, error) {
	if req.UserId == "" {
		return nil, errors.New("user id is empty")
	}

	return s.repo.SelectUsersByPreferences(ctx, req)
}
