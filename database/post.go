package database

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Post the structure thats represents the post data
//   saved values must start with a capitalized letter
type Post struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Title   string
	Body    string
	Created time.Time
}

// WritePost writes a post in the database
func WritePost(title, body string) error {
	posts := db.C("posts")
	post := Post{
		Title:   title,
		Body:    body,
		Created: time.Now(),
	}
	log.Print("trying to save: " + post.Title + " / " + post.Body)
	err := posts.Insert(post)
	checkError(err)
	return err
}

// ReadAllPosts read all the posts in the database
func ReadAllPosts() ([]Post, error) {
	var results []Post
	err := db.C("posts").Find(nil).All(&results)
	checkError(err)

	return results, err
}

// ReadPost read all the posts in the database
func ReadPost(id string) (*Post, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, &DatabaseError{
			time.Now(),
			"invalid ID",
		}
	}
	realID := bson.ObjectIdHex(id)
	var result Post
	err := db.C("posts").FindId(realID).One(&result)
	//err := c.Find(bson.M{"name": "Ale"}).One(&result)
	checkError(err)

	return &result, err
}
