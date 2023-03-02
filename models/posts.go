package models

import (
	"github.com/graphql-go/graphql"
)

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.Int},
			"userId": &graphql.Field{Type: graphql.Int},
			"title":  &graphql.Field{Type: graphql.String},
			"body":   &graphql.Field{Type: graphql.String},
		},
	},
)
