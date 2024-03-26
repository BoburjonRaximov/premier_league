package storage

import (
	"context"

	"github.com/BoburjonRaximov/premier_league/models"
)

type StorageI interface {
	Club() ClubesI
	Goal() GoalsI
	Match() MatchesI
	Result() ResultsI
}

type ClubesI interface {
	CreateClub(context.Context, models.CreateClub) (string, error)
	UpdateClub(context.Context, models.Club) (string, error)
	GetClub(context.Context, models.IdRequestClub) (models.Club, error)
	GetAllClub(context.Context, models.GetAllClubRequest) (models.GetAllClub, error)
	DeleteClub(context.Context, models.IdRequestClub) (string, error)
}

type GoalsI interface {
	CreateGoal(context.Context, models.CreateGoal) (string, error)
	GetGoal(context.Context, models.IdRequestGoal) (models.Goal, error)
	GetAllGoal(context.Context, models.GetAllGoalRequest) (models.GetAllGoal, error)
	DeleteGoal(context.Context, models.IdRequestGoal) (string, error)
}

type MatchesI interface {
	CreateMatch(context.Context, models.CreateMatch) (string, error)
	GetMatch(context.Context, models.IdRequestMatch) (models.Match, error)
	GetAllMatch(context.Context, models.GetAllMatchRequest) (models.GetAllMatch, error)
	DeleteMatch(context.Context, models.IdRequestMatch) (string, error)
}

type ResultsI interface {
	CreateResult(context.Context, models.CreateResult) (string, error)
	GetResult(context.Context, models.IdRequestResult) (models.Result, error)
	GetAllResult(context.Context, models.GetAllResultRequest) (models.GetAllResult, error)
	DeleteResult(context.Context, models.IdRequestResult) (string, error)
}
