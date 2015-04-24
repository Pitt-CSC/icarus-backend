// Package main provides ...
package main

import (
	. "github.com/Pitt-CSC/icarus-backend/routes"
	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	////
	// Handle routing
	////

	router := mux.NewRouter()
	api := router.
		PathPrefix("/api").
		Subrouter()

	// Talks
	talks := api.Path("/talks").Subrouter()
	talks.Methods("GET").HandlerFunc(TalkIndexRoute)

	talk := api.PathPrefix("/talks/{id}").Subrouter()
	talk.Methods("GET").HandlerFunc(TalkShowRoute)

	////
	// Add socket.io
	////

	io, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	io.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Emit("connection", nil)
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
	})

	io.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	router.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		io.ServeHTTP(w, r)
	})

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
