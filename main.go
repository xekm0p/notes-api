package main

import (
	"log"
	"net/http"
	"restapi/routes"
	"restapi/services"
	"restapi/utility"
)

func main() {
	var db = utility.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))

}
