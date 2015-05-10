// Package auth provides ...
package auth

import (
	"encoding/json"
	"github.com/Pitt-CSC/icarus-backend/models"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var db gorm.DB

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func InitializeDBConnection(dbconnection gorm.DB) {
	db = dbconnection
}

type githubAuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type githubUserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
}

func OAuthHandler(w http.ResponseWriter, request *http.Request) {
	// Get Github authentication information
	clientId := os.Getenv("ICARUS_GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("ICARUS_GITHUB_CLIENT_SECRET")

	// Get code from request
	code := request.FormValue("code")

	// Make POST request to Github's server to exchange code for auth token
	var url = "https://github.com/login/oauth/access_token"
	url += "?client_id=" + clientId
	url += "&client_secret=" + clientSecret
	url += "&code=" + code

	// Make a client for sending HTTP Requests
	client := &http.Client{}

	// Get the access token from Github
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}

	// Get the response and parse it for the JSON blob we're looking for
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}
	resp.Body.Close()
	auth := &githubAuthResponse{}
	if err := json.Unmarshal(body, &auth); err != nil {
		log.Print("There was an error using Unmarshal")
	}

	// Get the user ID from Github
	url = "https://api.github.com/user"
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "token "+auth.AccessToken)
	resp, err = client.Do(req)
	if err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}
	resp.Body.Close()
	githubUser := &githubUserResponse{}
	if err := json.Unmarshal(body, &githubUser); err != nil {
		// TODO: Do something to handle the error and return early, signifying that
		//       the authentication was unucessessful
	}

	var user models.User
	if err := db.Where(&models.User{GithubID: githubUser.ID}).First(&user).Error; err != nil {
		// User needs to be created
		createUser(githubUser)
		db.Where(&models.User{GithubID: githubUser.ID}).First(&user)
	}

	log.Printf("User name is: %s %s", user.FirstName, user.LastName)

	setSessionID(user.ID, w)

	http.Redirect(w, request, "http://localhost:4200", 301)
}

func GetAuthenticatedUser(request *http.Request) (models.User, error) {
	var user models.User

	// Get the decrypted ID from the cookie
	id, err := getSessionID(request)
	if err != nil {
		return user, err
	}

	// Get the user from the decrypted ID
	if err := db.Where(&models.User{ID: id}).First(&user).Error; err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

////
// Utility Functions
////

func createUser(githubUser *githubUserResponse) error {
	// Process their name
	nameArray := strings.Split(githubUser.Name, " ")
	firstName := nameArray[0]
	lastName := strings.Join(nameArray[1:len(nameArray)], " ")

	// Create the user object
	user := models.User{
		GithubID:  githubUser.ID,
		FirstName: firstName,
		LastName:  lastName,
		AvatarUrl: githubUser.AvatarUrl,
		Email:     githubUser.Email,
	}

	if db.NewRecord(user) {
		db.Create(&user)
		log.Printf("User #%d created", user.ID)
	}

	return nil
}

func setSessionID(id int, w http.ResponseWriter) error {
	value := map[string]int{
		"id": id,
	}
	encoded, err := cookieHandler.Encode("auth-token", value)
	if err != nil {
		log.Printf("Could not encode ID %d into cookie", id)
		return err
	}
	cookie := &http.Cookie{
		Name:  "auth-token",
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	return nil
}

func getSessionID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("auth-token")
	if err != nil {
		log.Print(err)
		log.Print("Could not get cookie from request")
		return 0, err
	}
	cookieValue := make(map[string]int)
	if err := cookieHandler.Decode("auth-token", cookie.Value, &cookieValue); err != nil {
		log.Print(err)
		return 0, nil
	}
	return cookieValue["id"], nil
}

////
// Error Handlers
////

func UnauthenticatedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}
