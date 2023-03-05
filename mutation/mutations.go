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
				"name":                 &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"username":             &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email":                &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address_street":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address_suite":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address_city":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address_zipcode":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address_geo_lat":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"address_geo_lng":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"phone":                &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"website":              &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"company_name":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"company_catch_phrase": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"company_bs":           &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return models.CreateUser(p)
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
				return models.CreatePost(p)
			},
		},
		"updateUser": &graphql.Field{
			Type:        models.UserType,
			Description: "Create a New User",
			Args: graphql.FieldConfigArgument{
				"id":                   &graphql.ArgumentConfig{Type: graphql.ID},
				"name":                 &graphql.ArgumentConfig{Type: graphql.String},
				"username":             &graphql.ArgumentConfig{Type: graphql.String},
				"email":                &graphql.ArgumentConfig{Type: graphql.String},
				"address_street":       &graphql.ArgumentConfig{Type: graphql.String},
				"address_suite":        &graphql.ArgumentConfig{Type: graphql.String},
				"address_city":         &graphql.ArgumentConfig{Type: graphql.String},
				"address_zipcode":      &graphql.ArgumentConfig{Type: graphql.String},
				"address_geo_lat":      &graphql.ArgumentConfig{Type: graphql.String},
				"address_geo_lng":      &graphql.ArgumentConfig{Type: graphql.String},
				"phone":                &graphql.ArgumentConfig{Type: graphql.String},
				"website":              &graphql.ArgumentConfig{Type: graphql.String},
				"company_name":         &graphql.ArgumentConfig{Type: graphql.String},
				"company_catch_phrase": &graphql.ArgumentConfig{Type: graphql.String},
				"company_bs":           &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return models.UpdateUser(p)
			},
		},
		"updatePost": &graphql.Field{
			Type:        models.PostType,
			Description: "Update a Post",
			Args: graphql.FieldConfigArgument{
				"id":     &graphql.ArgumentConfig{Type: graphql.ID},
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(string)
				userid, _ := p.Args["userid"].(string)
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
		"deleteUser": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete a Post",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := db.OpenConnection()
				id, _ := p.Args["id"].(string)
				row, err := db.Exec("DELETE FROM users WHERE id =$1 RETURNING id", id)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				rowsAffected, _ := row.RowsAffected()
				if rowsAffected == 0 {
					return nil, errors.New("user not found")
				}
				return true, nil
			},
		},
		"deletePost": &graphql.Field{
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
},
)
