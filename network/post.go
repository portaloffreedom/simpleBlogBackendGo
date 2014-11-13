package network

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/portaloffreedom/simpleBlogBackendGo/database"
)

// ReadAllPostsHandler is the handler for reading all the posts
func ReadAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, _ := database.ReadAllPosts()
	for i := 0; i < len(posts); i++ {
		fmt.Println("ID: " + string(posts[i].ID) + " - title: " + posts[i].Title + " - body: " + posts[i].Body)
	}
	writeJSON(w, posts)
}

// ReadPostHandler is the handler for reading a single post
func ReadPostHandler(w http.ResponseWriter, r *http.Request) {
	//"5464938b8fa9451e58353ae1"
	id := mux.Vars(r)["id"]

	//id := r.FormValue("id")
	post, _ := database.ReadPost(id)
	writeJSON(w, post)
}

// WritePostHandler is the handler for writing a post
func WritePostHandler(w http.ResponseWriter, r *http.Request) {
	database.WritePost("Ciao", "mamma zoccola")
	fmt.Fprint(w, "post written")
}
