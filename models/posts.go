package models

import (
	"github.com/graphql-go/graphql"
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userid"`
}

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.ID},
			"title":  &graphql.Field{Type: graphql.String},
			"body":   &graphql.Field{Type: graphql.String},
			"userid": &graphql.Field{Type: graphql.ID},
			// "user": &graphql.Field{
			// 	Type: UserType,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		db := db.OpenConnection()
			// 		parent, ok := p.Source.(*Post)
			// 		if !ok {
			// 			return nil, fmt.Errorf("invalid parent object")
			// 		}
			// 		if parent.UserId == 0 {
			// 			return nil, fmt.Errorf("invalid parent user ID")
			// 		}
			// 		user := &User{}
			// 		row := db.QueryRow("SELECT id, name, username, email FROM users WHERE id = $1", parent.UserId)
			// 		defer db.Close()
			// 		err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Email)
			// 		if err != nil {
			// 			return nil, err
			// 		}
			// 		return user, nil
			// 	},
			// },
		},
	},
)
