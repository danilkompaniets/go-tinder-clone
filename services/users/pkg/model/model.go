package model

import "time"

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
	UserId          string `json:"user_id" binding:"required"`
	PreferredGender string `json:"preferred_gender" binding:"required"`
	AgeMin          int    `json:"age_min" binding:"required"`
	AgeMax          int    `json:"age_max" binding:"required"`
	CityOnly        string `json:"city_only" binding:"required"`
}

type UpdatePreferenceRequest struct {
	ID              string  `json:"id"`
	PreferredGender *string `json:"preferred_gender" binding:"required"`
	AgeMin          *int    `json:"age_min" binding:"required"`
	AgeMax          *int    `json:"age_max" binding:"required"`
	CityOnly        *string `json:"city_only" binding:"required"`
}
