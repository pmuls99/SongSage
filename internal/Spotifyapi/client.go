package internalspotify

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func GetSpotifyBearerToken() (*spotify.Client, error) {
	godotenv.Load("local.env")
	clientId := os.Getenv("SPOTIFY_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	contentType := "application/x-www-form-urlencoded"
	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBuffer([]byte(form.Encode())))
	if err != nil {
		fmt.Println("Error sending the client authorization request to spotify")
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	authString := clientId + ":" + clientSecret
	stringEnc := base64.StdEncoding.EncodeToString([]byte(authString))
	authToken := "Basic " + stringEnc
	req.Header.Add("Authorization", authToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending the request to the spotify auth services")
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body :", err)
		return nil, err
	}
	// converting byte body to struct
	var authRespBody oauth2.Token
	// change the error name here
	eror := json.Unmarshal(respBody, &authRespBody)
	if eror != nil {
		fmt.Println("error unmarshaling the response body :", eror)
		return nil, eror
	}

	authenticator := spotifyauth.New()
	ctx := context.Background()

	// var oauthTokenBody oauth2.Token
	// json.NewDecoder(respBody).Decode(&oauthTokenBody)

	httpClient := authenticator.Client(ctx, &authRespBody)

	spotifyClient := spotify.New(httpClient)

	return spotifyClient, nil
	// return &authRespBody, nil
}
