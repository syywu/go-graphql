package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

func CreateCommentsTable() {
	db := OpenConnection()
	const createCommentTable = `
		CREATE TABLE IF NOT EXISTS comments(
		id SERIAL PRIMARY KEY,
		postId INT NOT NULL REFERENCES posts(id),
		name TEXT NOT NULL,
		email VARCHAR(255) NOT NULL,
		body TEXT
	);
	`
	_, err := db.Exec(createCommentTable)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func populateComments() []Comment {
	comment := Comment{ID: 1, PostID: 2, Name: "Sue", Email: "sue@gmail.com", Body: "first comment"}
	var comments []Comment
	comments = append(comments, comment)
	return comments
}

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"PostId": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Email": &graphql.Field{
				Type: graphql.String,
			},
			"Body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
