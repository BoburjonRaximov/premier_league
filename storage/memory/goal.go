package memory

import (
	"context"
	"fmt"

	"github.com/BoburjonRaximov/premier_league/models"
	"github.com/BoburjonRaximov/premier_league/pkg/helper"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type goalRepo struct {
	db *pgxpool.Pool
}

func NewGoalRepo(db *pgxpool.Pool) *goalRepo {
	return &goalRepo{db: db}
}

func (b *goalRepo) CreateGoal(ctx context.Context, req models.CreateGoal) (string, error) {
	fmt.Println("goal create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		goals(
			id,
			match_id,
			player_name,
			time) 
	VALUES($1,$2,$3,$4)`
	_, err := b.db.Exec(ctx, query,
		id,
		req.MatchId,
		req.PlayerName,
		req.Time,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (b *goalRepo) GetGoal(ctx context.Context, req models.IdRequestGoal) (models.Goal, error) {
	query := `
	SELECT
		id,
		match_id,
		player_name,
		time
	FROM 
		goals
	WHERE 
		id = $1`
	resp := b.db.QueryRow(ctx, query,
		req.Id,
	)
	var goal models.Goal
	err := resp.Scan(
		&goal.Id,
		&goal.MatchId,
		&goal.PlayerName,
		&goal.Time,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return goal, nil
}

func (b *goalRepo) GetAllGoal(ctx context.Context, req models.GetAllGoalRequest) (resp models.GetAllGoal, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = " WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		id,
		match_id,
		player_name,
		time
	FROM 
		goals
	`
	if req.Search != "" {
		filter += ` AND name ILIKE '%' || @search || '%' `
		params["search"] = req.Search
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := b.db.Query(ctx, q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		var goal models.Goal
		err := rows.Scan(
			&goal.Id,
			&goal.Id,
			&goal.Id,
			&goal.Id,
		)
		if err != nil {
			return resp, err
		}
		resp.Goals = append(resp.Goals, goal)
		resp.Count = len(resp.Goals)
	}
	return resp, nil
}

func (b *goalRepo) DeleteGoal(ctx context.Context, req models.IdRequestGoal) (string, error) {
	query := `
	DELETE FROM 
		goals
	WHERE 
		id=$1 `
	resp, err := b.db.Exec(ctx, query,
		req.Id,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "rows affected error", pgx.ErrNoRows
	}

	return "deleted", nil
}
