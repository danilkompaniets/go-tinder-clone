package model

import (
	"mime/multipart"
	"time"
)

type UpdateUserRequest struct {
	ID        string     `json:"id"` // обязателен
	Username  *string    `json:"username,omitempty"`
	FirstName *string    `json:"first_name,omitempty"`
	Bio       *string    `json:"bio,omitempty"`
	Gender    *string    `json:"gender,omitempty"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	City      *string    `json:"city,omitempty"`
	AvatarURL *string    `json:"avatar_url,omitempty"`
}

type User struct {
	ID        string    `json:"id"` // UUID в виде строки
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName *string   `json:"first_name,omitempty"`
	Bio       *string   `json:"bio,omitempty"`
	Gender    *string   `json:"gender,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	City      string    `json:"city"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePreferenceRequest struct {
	UserId          string    `json:"user_id" binding:"required"`
	PreferredGender string    `json:"preferred_gender" binding:"required"`
	AgeMin          int       `json:"age_min" binding:"required"`
	AgeMax          int       `json:"age_max" binding:"required"`
	Position        *Position `json:"position"`
}

type UserPreferences struct {
	ID              string    `json:"id"`
	UserId          *string   `json:"user_id"`
	PreferredGender *string   `json:"preferred_gender"`
	AgeMin          *int      `json:"age_min"`
	AgeMax          *int      `json:"age_max"`
	Position        *Position `json:"position"`
}

type Position struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius int     `json:"radius"` // in  meters
}

type GetUsersByPreferencesRequest struct {
	UserId          string    `json:"user_id" binding:"required"`
	PreferredGender string    `json:"preferred_gender" binding:"required"`
	AgeMin          int       `json:"age_min" binding:"required"`
	AgeMax          int       `json:"age_max" binding:"required"`
	Position        *Position `json:"position"`

	Limit  int `json:"limit" binding:"required"`
	Offset int `json:"offset" binding:"required"`
}

type AddUserPhotoRequest struct {
	Photo    *multipart.FileHeader `form:"photo" binding:"required"`
	Position *int                  `form:"position"`
}
