package mutation

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
	"github.com/syywu/go-graphql/models"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        models.UserType,
			Description: "Create a New User",
			Args: graphql.FieldConfigArgument{
				"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(models.AddressType)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				userid, _ := p.Args["userid"].(int)
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				var post models.Post
				_, err := db.Exec("INSERT INTO posts (userid, title, body) VALUES ($1, $2, $3) RETURNING id", userid, title, body)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				return post, nil
			},
		},
		"createPost": &graphql.Field{
			Type:        models.PostType,
			Description: "Create a New Post",
			Args: graphql.FieldConfigArgument{
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				userid, _ := p.Args["userid"].(int)
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				var post models.Post
				_, err := db.Exec("INSERT INTO posts (userid, title, body) VALUES ($1, $2, $3) RETURNING id", userid, title, body)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				return post, nil
			},
		},
		"update": &graphql.Field{
			Type:        models.PostType,
			Description: "Update a Post",
			Args: graphql.FieldConfigArgument{
				"id":     &graphql.ArgumentConfig{Type: graphql.Int},
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(int)
				userid, _ := p.Args["userid"].(int)
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				row, err := db.Exec("UPDATE posts SET userid = $1, title = $2, body = $3 WHERE id = $4 RETURNING *", userid, title, body, id)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				rowsAffected, _ := row.RowsAffected()
				if rowsAffected == 0 {
					return nil, errors.New("post not found")
				}
				return row, nil
			},
		},
		"delete": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete a Post",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(int)
				row, err := db.Exec("DELETE FROM posts WHERE id =$1 RETURNING id", id)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				rowsAffected, _ := row.RowsAffected()
				if rowsAffected == 0 {
					return nil, errors.New("post not found")
				}
				return true, nil
			},
		},
	},
})
