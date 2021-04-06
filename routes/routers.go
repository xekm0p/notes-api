package routes

import (
	"restapi/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/notes", services.GetAllNotes).Methods("GET")
	router.HandleFunc("/notes", services.CreateNote).Methods("POST")
	return router
}
