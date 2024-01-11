package main

import (
	routes "Guruprasad/Routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.CalorieTrackingRoutes(r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
