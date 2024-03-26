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

type resultRepo struct {
	db *pgxpool.Pool
}

func NewResultRepo(db *pgxpool.Pool) *resultRepo {
	return &resultRepo{db: db}
}

func (res *resultRepo) CreateResult(ctx context.Context, req models.CreateResult) (string, error) {
	fmt.Println("result create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		results(
			id,
			position,
			club_name,
			played,
			won,
			draw,
			lost,
			points) 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := res.db.Exec(ctx, query,
		id,
		req.Position,
		req.ClubNane,
		req.Played,
		req.Won,
		req.Draw,
		req.Lost,
		req.Point,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (res *resultRepo) GetResult(ctx context.Context, req models.IdRequestResult) (models.Result, error) {
	query := `
	SELECT
		position,
		club_name,
		played,
		won,
		draw,
		lost,
		points
	FROM 
		results
	WHERE 
		id = $1`
	resp := res.db.QueryRow(ctx, query,
		req.Id,
	)
	var result models.Result
	err := resp.Scan(
		&result.Id,
		&result.Position,
		&result.ClubNane,
		&result.Played,
		&result.Won,
		&result.Draw,
		&result.Lost,
		&result.Point,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return result, nil
}

func (b *resultRepo) GetAllResult(ctx context.Context, req models.GetAllResultRequest) (resp models.GetAllResult, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = " WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		position,
		club_name,
		played,
		won,
		draw,
		lost,
		points	
	FROM 
		results
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
		var result models.Result
		err := rows.Scan(
			&result.Id,
			&result.Position,
			&result.ClubNane,
			&result.Played,
			&result.Won,
			&result.Draw,
			&result.Lost,
			&result.Point,
		)
		if err != nil {
			return resp, err
		}
		resp.Results = append(resp.Results, result)
		resp.Count = len(resp.Results)
	}
	return resp, nil
}

func (b *resultRepo) DeleteResult(ctx context.Context, req models.IdRequestResult) (string, error) {
	query := `
	DELETE FROM 
		results
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
