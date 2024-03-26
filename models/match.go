package models

type CreateMatch struct {
	Id            string `json:"id"`
	ResultId      string `json:"result_id"`
	MatchDate       string `json:"match_id"`
	Attendance    int    `json:"attendance"`
	Referee       string `json:"referee"`
	ClubId        string `json:"club_id"`
	AgainstClubId string `json:"against_club_id"`
}

type Match struct {
	Id            string `json:"id"`
	ResultId      string `json:"result_id"`
	MatchDate       string `json:"match_id"`
	Attendance    int    `json:"attendance"`
	Referee       string `json:"referee"`
	ClubId        string `json:"club_id"`
	AgainstClubId string `json:"against_club_id"`
	CreatedAt     string `json:"created_at"`
}

type GetAllMatchRequest struct {
	Page   int
	Limit  int
	Search string
}

type GetAllMatch struct {
	Matches []Match
	Count   int
}

type IdRequestMatch struct {
	Id string
}
