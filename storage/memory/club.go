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

type clubRepo struct {
	db *pgxpool.Pool
}

func NewClubRepo(db *pgxpool.Pool) *clubRepo {
	return &clubRepo{db: db}
}

func (b *clubRepo) CreateClub(ctx context.Context, req models.CreateClub) (string, error) {
	fmt.Println("club create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		clubs(
			id,
			name,
			stadium,
			established) 
	VALUES($1,$2,$3,$4)`
	_, err := b.db.Exec(ctx, query,
		id,
		req.Name,
		req.Stadium,
		req.Established,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (b *clubRepo) GetClub(ctx context.Context, req models.IdRequestClub) (models.Club, error) {
	query := `
	SELECT
		id,
		name,
		stadium,
		established
	FROM 
		clubs
	WHERE 
		id = $1`
	resp := b.db.QueryRow(ctx, query,
		req.Id,
	)
	var club models.Club
	err := resp.Scan(
		&club.Id,
		&club.Name,
		&club.Stadium,
		&club.Established,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return club, nil
}

func (b *clubRepo) UpdateClub(ctx context.Context, req models.Club) (string, error) {
	query := `
	UPDATE 
		clubs
	SET 
		name=$2,
		stadium=$3,
		established=$4
	WHERE 
		id=$1`
	resp, err := b.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Stadium,
		req.Established,
	)
	if err != nil {
		return "warning", err
	}
	if resp.RowsAffected() == 0 {
		return "error row", pgx.ErrNoRows
	}
	return "OK", nil
}

func (b *clubRepo) GetAllClub(ctx context.Context, req models.GetAllClubRequest) (resp models.GetAllClub, err error) {
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
		name,
		stadium,
		established
	FROM 
		clubs
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
		var club models.Club
		err := rows.Scan(
			&club.Id,
			&club.Name,
			&club.Stadium,
			&club.Established,
		)
		if err != nil {
			return resp, err
		}
		resp.Clubs = append(resp.Clubs, club)
		resp.Count = len(resp.Clubs)
	}
	return resp, nil
}

func (b *clubRepo) DeleteClub(ctx context.Context, req models.IdRequestClub) (string, error) {
	query := `
	DELETE FROM 
		clubs
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
