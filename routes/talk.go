// Package routes provides ...
package routes

import (
	"encoding/json"
	"github.com/Pitt-CSC/icarus-backend/models"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func TalkIndexRoute(w http.ResponseWriter, r *http.Request) {
	//talk1 := models.Talk{ID: 1, Title: "This is the first talk."}
	//talk2 := models.Talk{ID: 2, Title: "This is the second talk."}
	var talks []models.Talk
	if err := db.Find(&talks).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	talkarray := models.TalkCollection{Data: talks}
	sendJSON(w, talkarray)
}

func TalkNewRoute(w http.ResponseWriter, r *http.Request) {
	var talk models.Talk
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &talk); err != nil {
		sendUnprocessableEntity(w, err)
	}
	if db.NewRecord(talk) {
		db.Create(&talk)
		log.Printf("Talk #%d created", talk.ID)
	}
	w.WriteHeader(http.StatusCreated)
	sendJSON(w, talk)
}

func TalkShowRoute(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var talk models.Talk
	if err := db.Where("id = ?", id).First(&talk).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	talkWrapper := models.TalkResource{Data: talk}
	sendJSON(w, talkWrapper)
}

func TalkDeleteRoute(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	talk := models.Talk{ID: id}
	db.Delete(&talk)
	w.WriteHeader(http.StatusNoContent)
}
