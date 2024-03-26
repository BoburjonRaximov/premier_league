package models

type CreateResult struct {
	Id       string `json:"id"`
	Position int    `json:"position"`
	ClubNane string `json:"club_name"`
	Played   int    `json:"played"`
	Won      int    `json:"won"`
	Draw     int    `json:"draw"`
	Lost     int    `json:"lost"`
	Point    int    `json:"points"`
}

type Result struct {
	Id        string `json:"id"`
	Position  int    `json:"position"`
	ClubNane  string `json:"club_name"`
	Played    int    `json:"played"`
	Won       int    `json:"won"`
	Draw      int    `json:"draw"`
	Lost      int    `json:"lost"`
	Point     int    `json:"points"`
	CreatedAt string `json:"created_at"`
}

type GetAllResultRequest struct {
	Page   int
	Limit  int
	Search string
}

type GetAllResult struct {
	Results []Result
	Count   int
}

type IdRequestResult struct {
	Id string
}
