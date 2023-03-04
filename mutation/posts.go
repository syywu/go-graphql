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
				userid, _ := p.Args["userid"].(string)
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
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				userid, _ := p.Args["userid"].(string)
				var post models.Post
				_, err := db.Exec("INSERT INTO posts (title, body, userid) VALUES ($1, $2, $3) RETURNING id", title, body, userid)
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
				"id":     &graphql.ArgumentConfig{Type: graphql.ID},
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(string)
				userid, _ := p.Args["userid"].(int)
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				row, err := db.Exec("UPDATE posts SET title = $1, body = $2, userid = $3 WHERE id = $4 RETURNING *", title, body, id, userid)
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
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(string)
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
