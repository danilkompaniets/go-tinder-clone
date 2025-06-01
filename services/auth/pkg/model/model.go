package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterUserRequest struct {
	Username  string    `json:"username"`
	Password  string    `json:"password" binding:"required,min=6"`
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

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type GetUserRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
