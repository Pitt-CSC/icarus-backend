// Package routes provides ...
package routes

import (
	"encoding/json"
	. "github.com/Pitt-CSC/icarus-backend/models"
	"github.com/gorilla/mux"
	"net/http"
)

func sendJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(obj)
}

func TalkIndexRoute(w http.ResponseWriter, r *http.Request) {
	talk1 := Talk{Id: "1", Title: "This is a talk"}
	talk2 := Talk{Id: "2", Title: "This is a talk"}
	talkarray := TalkCollection{Data: []Talk{talk1, talk2}}
	sendJSON(w, talkarray)
}

func TalkShowRoute(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	talk := Talk{Id: id, Title: "This is a talk"}
	talkWrapper := TalkResource{Data: talk}
	sendJSON(w, talkWrapper)
}
