// Package main provides ...
package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
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

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		io.ServeHTTP(w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
