package routes

import (
	"log"
	"net/http"

	"github.com/maitecr/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.GetPersonalidades)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
