package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type NetworkError struct {
	When time.Time
	What string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// StartServer stars an http server
func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/read", ReadAllPostsHandler)
	r.HandleFunc("/read/{id:[a-z0-9]+}", ReadPostHandler)
	r.HandleFunc("/write", WritePostHandler)
	http.Handle("/", r)

	http.ListenAndServe("localhost:4000", nil)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	jsonResponse, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(jsonResponse)
}

// HomeHandler is the handler for the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
