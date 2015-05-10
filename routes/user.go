// Package routes
package routes

import (
	"github.com/Pitt-CSC/icarus-backend/auth"
	"net/http"
)

func GetAuthenticatedUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetAuthenticatedUser(r)
	if err != nil {
		auth.UnauthenticatedHandler(w, r)
		return
	}
	sendJSON(w, user)
}
