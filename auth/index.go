// Package auth provides ...
package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GithubAuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func LoginHandler(w http.ResponseWriter, request *http.Request) {
	// Get Github authentication information
	clientId := os.Getenv("ICARUS_GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("ICARUS_GITHUB_CLIENT_SECRET")

	// Get code from request
	code := request.FormValue("code")
	log.Printf("Authentication code %s", code)

	// Make POST request to Github's server to exchange code for auth token
	var url = "https://github.com/login/oauth/access_token"
	url += "?client_id=" + clientId
	url += "&client_secret=" + clientSecret
	url += "&code=" + code

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}
	defer resp.Body.Close()

	// Get the response and parse it for the JSON blob we're looking for
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}

	auth := &GithubAuthResponse{}
	if err := json.Unmarshal(body, &auth); err != nil {
		log.Print("There was an error using Unmarshal")
	}

	log.Printf("Access token is %s", auth.AccessToken)

	http.Redirect(w, request, "http://localhost:4200", 301)
}
