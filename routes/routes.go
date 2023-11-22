package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maitecr/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetById).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
