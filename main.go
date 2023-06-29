package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	internalspotify "github.com/pmuls99/likeSongs/internal/Spotifyapi"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getAuthToken", authTokenHandler)
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
