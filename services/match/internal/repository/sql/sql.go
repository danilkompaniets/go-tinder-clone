package sqlRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/danilkompanites/tinder-clone/services/match/pkg/model"
)

type SqlRepository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{db: db}
}

func (r *SqlRepository) InsertMatch(ctx context.Context, request model.Match) (*model.Match, error) {
	stmt := `
		INSERT INTO match (from_id, to_id, from_decision) VALUES ($1, $2, $3) 
		ON CONFLICT DO UPDATE SET to_decision = $3
		RETURNING from_id, to_id, from_decision, to_decision
	`

	var res model.Match

	err := r.db.QueryRowContext(ctx, stmt, request.FromId, request.ToId, request.FromDecision).Scan(&res.FromId, &res.ToId, &res.FromDecision, &res.ToDecision)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SqlRepository) SelectAllDecisionsByUserId(ctx context.Context, userId string) ([]*model.Match, error) {
	query := `SELECT * FROM match WHERE from_decision = $1`

	rows, err := r.db.QueryContext(ctx, query, userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var res []*model.Match
	for rows.Next() {
		var m model.Match
		err := rows.Scan(&m.FromId, &m.ToId, &m.FromDecision, &m.ToDecision)
		if err != nil {
			return nil, err
		}
		res = append(res, &m)
	}

	return res, nil
}
