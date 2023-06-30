package model

import "github.com/zmb3/spotify/v2"

type SearchRequest struct {
	Search     string             `json:"search"`
	SearchType spotify.SearchType `json:"search_type"`
}
