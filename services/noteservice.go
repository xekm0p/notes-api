package services

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/models"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var notes = models.GetNotes()
	var resp Response
	err := dbconn.Find(&notes).Error
	if err == nil {
		log.Println(notes)
		resp.Data = notes
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note = models.GetNote()
	_ = json.NewDecoder(r.Body).Decode(&note)
	err := dbconn.Create(&note).Error
	if err == nil {

		json.NewEncoder(w).Encode(note)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}
