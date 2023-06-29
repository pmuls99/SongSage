package internalspotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	model "github.com/pmuls99/likeSongs/model/spotify"
)

func GetSpotifyBearerToken() (*model.AuthTokenResponse, error) {
	godotenv.Load("local.env")
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
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
	var authRespBody model.AuthTokenResponse
	// change the error name here
	eror := json.Unmarshal(respBody, &authRespBody)
	if eror != nil {
		fmt.Println("error unmarshaling the response body :", eror)
		return nil, eror
	}

	return &authRespBody, nil
}
