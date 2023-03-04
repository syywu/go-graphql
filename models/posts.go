package models

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
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
			"id":    &graphql.Field{Type: graphql.Int},
			"title": &graphql.Field{Type: graphql.String},
			"body":  &graphql.Field{Type: graphql.String},
			"userId": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if parent, ok := p.Source.(*Post); ok {
						user := &User{}
						db := db.OpenConnection()
						err := db.QueryRow("SELECT * FROM users WHERE id = $1", parent.UserId).Scan(&user.Name, &user.Username, &user.Email, &user.Address, &user.Phone, &user.Website, &user.Company)
						if err != nil {
							log.Fatal(err)
						}
						defer db.Close()
						return user, nil
					}
					return nil, nil
				},
			},
		},
	},
)
