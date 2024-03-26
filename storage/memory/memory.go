package memory

import (
	"context"
	"fmt"

	"github.com/BoburjonRaximov/premier_league/config"
	"github.com/BoburjonRaximov/premier_league/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db     *pgxpool.Pool
	club   *clubRepo
	goal   *goalRepo
	match  *matchRepo
	result *resultRepo
}

func (c *strg) Club() storage.ClubesI {
	if c.club == nil {
		c.club = NewClubRepo(c.db)
	}
	return c.club
}

func (c *strg) Goal() storage.GoalsI {
	if c.goal == nil {
		c.goal = NewGoalRepo(c.db)
	}
	return c.goal
}

func (c *strg) Match() storage.MatchesI {
	if c.match == nil {
		c.match = NewMatchRepo(c.db)
	}
	return c.match
}

func (c *strg) Result() storage.ResultsI {
	if c.result == nil {
		c.result = NewResultRepo(c.db)
	}
	return c.result
}

func NewStorage(ctx context.Context, cfg config.ConfigPostgres) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}
	return &strg{
		db: pool,
	}, nil
}
