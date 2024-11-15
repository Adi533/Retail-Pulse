package main

import (
	"log"
	"net/http"
	"retailpulse/job"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/api/submit/", job.SubmitJobHandler).Methods("POST")
	r.HandleFunc("/api/status", job.GetJobInfoHandler).Methods("GET")

	// start server
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
