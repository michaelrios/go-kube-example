package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"os"
	log "github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	logger := log.Logger{
		Out:       os.Stdout,
		Formatter: &log.JSONFormatter{},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler)


	router.HandleFunc("/", addNameHandler).
		Methods("POST")

	server := http.Server{
		Addr: ":80",
		Handler: router,
	}
	defer server.Close()

	logger.Fatal(server.ListenAndServe())
}

// WelcomeMessage is a generic welcome message
type WelcomeMessage struct {
	Message string `json:"message"`
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get name from the DB

	message := WelcomeMessage{fmt.Sprintf("Hello %s!", "Tater Totten")}
	jsonReturn, _ := json.Marshal(message)
	w.Write(jsonReturn)
}

func addNameHandler(w http.ResponseWriter, r *http.Request) {
	// Add name to DB

	w.WriteHeader(http.StatusAccepted)
}