package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type NetworkError struct {
	status int
	Code   string
	When   time.Time
	What   string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// StartServer stars an http server
func StartServer(bindAddress string) {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/read", ReadAllPostsHandler)
	r.HandleFunc("/read/{id:[a-z0-9]+}", ReadPostHandler)
	r.HandleFunc("/write", WritePostHandler)
	http.Handle("/", r)

	fmt.Printf("listening on %s...", bindAddress)
	err := http.ListenAndServe(bindAddress, nil)
	if err != nil {
		log.Fatal("Error listening to the port: " + err.Error())
	}
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	jsonResponse, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(jsonResponse)
}

func writeError(w http.ResponseWriter, err *NetworkError) {
	if err == nil {
		err = &NetworkError{
			http.StatusInternalServerError,
			"generic_error",
			time.Now(),
			"Generic Error Occoured",
		}
	}
	jsonResponse, _ := json.Marshal(err)
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(err.status)
	w.Write(jsonResponse)
}

// HomeHandler is the handler for the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
