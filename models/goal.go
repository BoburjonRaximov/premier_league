package models

type CreateGoal struct {
	Id         string `json:"id"`
	MatchId    string `json:"match_id"`
	PlayerName string `json:"player_name"`
	Time       string `json:"time"`
}

type Goal struct {
	Id         string `json:"id"`
	MatchId    string `json:"match_id"`
	PlayerName string `json:"player_name"`
	Time       string `json:"time"`
	CreatedAt  string `json:"created_at"`
}

type GetAllGoalRequest struct {
	Page   int
	Limit  int
	Search string
}

type GetAllGoal struct {
	Goals []Goal
	Count int
}

type IdRequestGoal struct {
	Id string
}
