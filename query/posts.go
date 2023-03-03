package query

import (
	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
	"github.com/syywu/go-graphql/models"
)

// define our schema- define what fields to return to us for when making queries
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type:        models.PostType,
			Description: "Get Post by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(int)
				row := db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
				defer db.Close()

				var post models.Post
				err := row.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
				if err != nil {
					return nil, err
				}
				return post, nil
			},
		},
		"posts": &graphql.Field{
			Type:        graphql.NewList(models.PostType),
			Description: "Get All Posts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				rows, err := db.Query("SELECT * FROM posts")
				if err != nil {
					return nil, err
				}
				defer db.Close()
				defer rows.Close()
				posts := []models.Post{}

				for rows.Next() {
					var post models.Post
					err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
					if err != nil {
						return nil, err
					}
					posts = append(posts, post)
				}
				return posts, nil
			},
		},
	},
})
