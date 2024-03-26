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

type matchRepo struct {
	db *pgxpool.Pool
}

func NewMatchRepo(db *pgxpool.Pool) *matchRepo {
	return &matchRepo{db: db}
}

func (b *matchRepo) CreateMatch(ctx context.Context, req models.CreateMatch) (string, error) {
	fmt.Println("club create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		matches(
			id,
			result_id,
			match_date,
			attendance,
			referee,
			club_id,
			against_club_id) 
	VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := b.db.Exec(ctx, query,
		id,
		req.ResultId,
		req.MatchDate,
		req.Attendance,
		req.Referee,
		req.ClubId,
		req.AgainstClubId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (b *matchRepo) GetMatch(ctx context.Context, req models.IdRequestMatch) (models.Match, error) {
	query := `
	SELECT
		id,
		result_id,
		match_date,
		attendance,
		referee,
		club_id,
		against_club_id
	FROM 
		matches
	WHERE 
		id = $1`
	resp := b.db.QueryRow(ctx, query,
		req.Id,
	)
	var match models.Match
	err := resp.Scan(
		&match.Id,
		&match.ResultId,
		&match.MatchDate,
		&match.Attendance,
		&match.Referee,
		&match.ClubId,
		&match.AgainstClubId,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return match, nil
}

func (b *matchRepo) GetAllMatch(ctx context.Context, req models.GetAllMatchRequest) (resp models.GetAllMatch, err error) {
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
		result_id,
		match_date,
		attendance,
		referee,
		club_id,
		against_club_id
	FROM 
		matches
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
		var match models.Match
		err := rows.Scan(
			&match.Id,
			&match.Id,
			&match.Id,
			&match.Id,
		)
		if err != nil {
			return resp, err
		}
		resp.Matches = append(resp.Matches, match)
		resp.Count = len(resp.Matches)
	}
	return resp, nil
}
func (b *matchRepo) DeleteMatch(ctx context.Context, req models.IdRequestMatch) (string, error) {
	query := `
	DELETE FROM 
		matches
	WHERE 
		id=$1 `
	resp, err := b.db.Exec(ctx, query,
		req.Id,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "deleted", nil
}
