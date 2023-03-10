package models

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
)

type Post struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Userid int    `json:"userid"`
}

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if post, ok := p.Source.(*Post); ok {
						return post.ID, nil
					}

					return nil, nil
				}},
			"title":  &graphql.Field{Type: graphql.String},
			"body":   &graphql.Field{Type: graphql.String},
			"userid": &graphql.Field{Type: graphql.ID},
			"user": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.OpenConnection()
					parent, ok := p.Source.(*Post)
					if !ok {
						return nil, fmt.Errorf("invalid parent object")
					}
					if parent.Userid == 0 {
						return nil, fmt.Errorf("invalid parent user ID")
					}
					user := &User{}
					err := db.QueryRow("SELECT id, name, username, email FROM users WHERE id = $1", parent.Userid).Scan(&user.ID, &user.Name, &user.Username, &user.Email)
					defer db.Close()
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
		},
	},
)
