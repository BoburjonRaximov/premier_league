package models

type CreateClub struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Stadium     string `json:"stadium"`
	Established int    `json:"established"`
}

type Club struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Stadium     string `json:"stadium"`
	Established int    `json:"established"`
	CreatedAt   string `json:"created_at"`
}

type GetAllClubRequest struct {
	Page   int
	Limit  int
	Search string
}

type GetAllClub struct {
	Clubs []Club
	Count int
}

type IdRequestClub struct {
	Id string
}
