package model

type SearchRequest struct {
	Search   string `json:"search"`
	SearchBy string `json:"searchBy"`
}
