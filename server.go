// Package main provides ...
package main

import (
	"github.com/Pitt-CSC/icarus-backend/auth"
	"github.com/Pitt-CSC/icarus-backend/models"
	"github.com/Pitt-CSC/icarus-backend/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	////
	// Database Connection
	//
	// See migrate.go for migrations
	////

	db, err := gorm.Open("sqlite3", "tmp/gorm.db")
	if err != nil {
		panic(err)
	}
	db.DB()

	////
	// Handle routing
	////

	routes.InitializeDBConnection(db)
	models.InitializeDBConnection(db)
	router := mux.NewRouter()
	api := router.
		PathPrefix("/api").
		Subrouter()

	// Authentiation
	oauth := api.Path("/oauth").Subrouter()
	oauth.Methods("GET").HandlerFunc(auth.OAuthHandler)

	// Talks
	talks := api.Path("/talks").Subrouter()
	talks.Methods("GET").HandlerFunc(routes.TalkIndexRoute)
	talks.Methods("POST").HandlerFunc(routes.TalkNewRoute)

	talk := api.PathPrefix("/talks/{id}").Subrouter()
	talk.Methods("GET").HandlerFunc(routes.TalkShowRoute)
	talk.Methods("DELETE").HandlerFunc(routes.TalkDeleteRoute)

	////
	// Add websockets
	////

	h := NewHub()
	go h.run()
	router.Handle("/socket/", WsHandler{h: h})

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
