package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleResquets() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
