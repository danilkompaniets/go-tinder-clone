package sql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/danilkompanites/tinder-clone/services/auth/pkg/model"
	"github.com/google/uuid"
)

type Repository struct {
	DB *sql.DB
}

func NewSqlRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) InsertUser(ctx context.Context, user *model.User) (string, error) {
	id := uuid.New().String()
	query := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var insertedID string
	err := r.DB.QueryRowContext(ctx, query, id, user.Email, user.Password).Scan(&insertedID)
	if err != nil {
		return "", err
	}
	return insertedID, nil
}

func (r *Repository) SelectUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, email, password, created_at, updated_at, is_active
		FROM users
		WHERE email = $1
	`
	var user model.User
	err := r.DB.QueryRowContext(ctx, query, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) SelectUserByRefreshToken(ctx context.Context, token string) (*model.User, error) {
	query := `
		SELECT u.id, u.email, u.password, u.created_at, u.updated_at, u.is_active
		FROM refresh_tokens rt
		JOIN users u ON u.id = rt.user_id
		WHERE rt.token = $1
	`
	var user model.User
	err := r.DB.QueryRowContext(ctx, query, token).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) InsertRefreshToken(ctx context.Context, userID string, token string) error {
	query := `INSERT INTO refresh_tokens (user_id, token) VALUES ($1, $2)`
	_, err := r.DB.ExecContext(ctx, query, userID, token)
	return err
}

func (r *Repository) DeleteRefreshToken(ctx context.Context, token string) error {
	query := `DELETE FROM refresh_tokens WHERE token = $1`
	_, err := r.DB.ExecContext(ctx, query, token)
	return err
}

func (r *Repository) DeleteAllRefreshTokensByUserID(ctx context.Context, userID string) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	_, err := r.DB.ExecContext(ctx, query, userID)
	return err
}
