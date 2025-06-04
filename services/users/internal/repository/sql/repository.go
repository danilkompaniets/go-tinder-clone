package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilkompanites/tinder-clone/services/users/pkg/model"
	"strings"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) InsertUser(ctx context.Context, req model.User) error {
	query := `
		INSERT INTO users (
			id, username, email, first_name, bio, gender, birth_date, city, avatar_url, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := repo.DB.ExecContext(ctx, query,
		req.ID,
		req.Username,
		req.Email,
		req.FirstName,
		req.Bio,
		req.Gender,
		req.BirthDate,
		req.City,
		req.AvatarURL,
		time.Now(),
		time.Now(),
	)

	return err
}

func (repo *Repository) SelectUserByID(ctx context.Context, id string) (*model.User, error) {
	query := `
		SELECT id, username, email, first_name, bio, gender, birth_date, city, avatar_url, created_at, updated_at
		FROM users WHERE id = $1 
	`

	var user model.User
	err := repo.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.Bio,
		&user.Gender,
		&user.BirthDate,
		&user.City,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) InsertUserPreference(ctx context.Context, req model.CreatePreferenceRequest) (string, error) {
	query := `
		INSERT INTO users_preferences (
			user_id, preferred_gender, age_min, age_max, city_only, position, radius
		) 
		VALUES (
			$1, $2, $3, $4, ST_MakePoint($5, $6)::geography, $7
		) RETURNING id
	`

	lon := req.Position.Lon
	lat := req.Position.Lat
	radius := req.Position.Radius

	var id string
	err := repo.DB.QueryRowContext(ctx, query,
		req.UserId,
		req.PreferredGender,
		req.AgeMin,
		req.AgeMax,
		lon,
		lat,
		radius,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (repo *Repository) DeleteUserPreference(ctx context.Context, id string) error {
	query := "DELETE FROM users_preferences WHERE id = $1"
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}

func (repo *Repository) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) error {
	var setParts []string
	var args []interface{}
	i := 1

	if req.Username != nil {
		setParts = append(setParts, fmt.Sprintf("username = $%d", i))
		args = append(args, *req.Username)
		i++
	}
	if req.FirstName != nil {
		setParts = append(setParts, fmt.Sprintf("first_name = $%d", i))
		args = append(args, *req.FirstName)
		i++
	}
	if req.Bio != nil {
		setParts = append(setParts, fmt.Sprintf("bio = $%d", i))
		args = append(args, *req.Bio)
		i++
	}
	if req.Gender != nil {
		setParts = append(setParts, fmt.Sprintf("gender = $%d", i))
		args = append(args, *req.Gender)
		i++
	}
	if req.BirthDate != nil {
		setParts = append(setParts, fmt.Sprintf("birth_date = $%d", i))
		args = append(args, *req.BirthDate)
		i++
	}
	if req.City != nil {
		setParts = append(setParts, fmt.Sprintf("city = $%d", i))
		args = append(args, *req.City)
		i++
	}
	if req.AvatarURL != nil {
		setParts = append(setParts, fmt.Sprintf("avatar_url = $%d", i))
		args = append(args, *req.AvatarURL)
		i++
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", i))
	args = append(args, time.Now())
	i++

	if len(setParts) == 0 {
		return errors.New("no fields to update")
	}

	query := `UPDATE users SET ` + strings.Join(setParts, ", ") + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, req.ID)

	_, err := repo.DB.ExecContext(ctx, query, args...)
	return err
}

func (repo *Repository) UpdateUserPreference(ctx context.Context, req model.UserPreferences) error {
	query := `UPDATE users_preferences SET `
	var args []interface{}
	i := 1

	if req.PreferredGender != nil {
		query += fmt.Sprintf("preferred_gender = $%d, ", i)
		args = append(args, *req.PreferredGender)
		i++
	}
	if req.AgeMin != nil {
		query += fmt.Sprintf("age_min = $%d, ", i)
		args = append(args, *req.AgeMin)
		i++
	}
	if req.AgeMax != nil {
		query += fmt.Sprintf("age_max = $%d, ", i)
		args = append(args, *req.AgeMax)
		i++
	}
	if req.Position != nil && req.Position.Lat != nil && req.Position.Lon != nil && req.Position.Radius != nil {
		query += fmt.Sprintf("position = ST_MakePoint($%d, $%d)::geography, ", i, i+1)
		args = append(args, *req.Position.Lon, *req.Position.Lat)
		i += 2
		query += fmt.Sprintf("radius = $%d, ", i)
		args = append(args, *req.Position.Radius)
		i++
	}

	if len(args) == 0 {
		return errors.New("no fields to update")
	}

	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, req.ID)

	_, err := repo.DB.ExecContext(ctx, query, args...)
	return err
}

func (repo *Repository) SelectUsersByPreferences(ctx context.Context, req *model.GetUsersByPreferencesRequest) ([]*model.User, error) {
	query := `
		SELECT *
		FROM users u
		JOIN users_preferences up ON up.user_id = u.id
		WHERE u.gender = $1
		  AND ST_DWithin(
			up.position, -- география пользователя
			ST_MakePoint($2, $3)::geography, -- координаты центра поиска (lon, lat)
			$4 -- радиус в метрах
		  )
		  AND DATE_PART('year', AGE(u.birth_date)) BETWEEN $5 AND $6
		LIMIT $7 OFFSET $8;	
	`

	rows, err := repo.DB.QueryContext(ctx, query,
		req.PreferredGender,
		req.Position.Lon,
		req.Position.Lat,
		req.AgeMin,
		req.AgeMax,
		req.Limit,
		req.Offset,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.User, 1)

	for rows.Next() {
		var usr model.User
		err := rows.Scan(
			&usr.ID,
			&usr.Username,
			&usr.Email,
			&usr.FirstName,
			&usr.Bio,
			&usr.Gender,
			&usr.BirthDate,
			&usr.City,
			&usr.AvatarURL,
			&usr.CreatedAt,
			&usr.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, &usr)
	}

	return res, nil
}
