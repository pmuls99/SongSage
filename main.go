package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	internalspotify "github.com/pmuls99/likeSongs/internal/Spotifyapi"
	model "github.com/pmuls99/likeSongs/model/search"
	spotify "github.com/zmb3/spotify/v2"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getAuthToken", authTokenHandler)
	router.HandleFunc("/getSearchResults", searchHandler)
	port := ":8000"
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println(" error listening and serving at the current port ", err)
	}
}

// getting the auth token from the spotify service

func authTokenHandler(w http.ResponseWriter, r *http.Request) {
	authRespBody, err := internalspotify.GetSpotifyBearerToken()
	if err != nil {
		fmt.Println("Error fetching the response body from the server side :", err)
	}
	// var resp http.Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authRespBody)
}

// Get Search Recommendations

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// we recieve a search request body read the body from json type to native lang
	var searchReqBody model.SearchRequest
	// fmt.Println(r.Body)
	// requestBody, err := io.ReadAll(r.Body)
	searchReqBody.Search = r.URL.Query().Get("search")
	searchType, err := strconv.Atoi(r.URL.Query().Get("searchType"))
	searchReqBody.SearchType = spotify.SearchType(searchType)
	if err != nil {
		fmt.Println("error fetching the search request body : ", err)
	}

	// fmt.Println(string(requestBody))
	// json.NewDecoder(r.Body).Decode(&searchReqBody)
	fmt.Println(searchReqBody)

	spotifyClient, err := internalspotify.GetSpotifyBearerToken()
	if err != nil {
		fmt.Println("Error fetching the spotify token")
	}

	// authenticator := authSpotify.NewAuthenticator("http://localhost:8000/", "streaming")

	// spotifyClient := authSpotify.Authenticator.NewClient(authenticator, &oauthTokenBody)
	var spotifySearchResult *spotify.SearchResult

	spotifySearchResult, err = spotifyClient.Search(r.Context(), searchReqBody.Search, searchReqBody.SearchType)
	if err != nil {
		fmt.Println("Possibly an error with the spotify client : ", err)
	}
	// json.Encoder(*spotifySearchResult).Encode(model.SearchResponse)
	// json.NewEncoder(w).Encode(&spotifySearchResult)
	json.NewEncoder(w).Encode(&spotifySearchResult)

}
