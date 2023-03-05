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
				db := db.OpenConnection()
				name, _ := p.Args["name"].(string)
				username, _ := p.Args["username"].(string)
				email, _ := p.Args["email"].(string)
				address_street, _ := p.Args["address_street"].(string)
				address_suite, _ := p.Args["address_suite"].(string)
				address_city, _ := p.Args["address_city"].(string)
				address_zipcode, _ := p.Args["address_zipcode"].(string)
				address_geo_lat, _ := p.Args["address_geo_lat"].(string)
				address_geo_lng, _ := p.Args["address_geo_lng"].(string)
				phone, _ := p.Args["phone"].(string)
				website, _ := p.Args["website"].(string)
				company_name, _ := p.Args["company_name"].(string)
				company_catch_phrase, _ := p.Args["company_catch_phrase"].(string)
				company_bs, _ := p.Args["company_bs"].(string)
				var post models.Post
				_, err := db.Exec("INSERT INTO users (name, username, email, address_street,address_suite, address_city, address_zipcode, address_geo_lat, address_geo_lng, phone, website, company_name, company_catch_phrase, company_bs) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id", name, username, email, address_street, address_suite, address_city, address_zipcode, address_geo_lat, address_geo_lng, phone, website, company_name, company_catch_phrase, company_bs)
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
				db := db.OpenConnection()
				id, _ := p.Args["id"].(string)
				name, _ := p.Args["name"].(string)
				username, _ := p.Args["username"].(string)
				email, _ := p.Args["email"].(string)
				address_street, _ := p.Args["address_street"].(string)
				address_suite, _ := p.Args["address_suite"].(string)
				address_city, _ := p.Args["address_city"].(string)
				address_zipcode, _ := p.Args["address_zipcode"].(string)
				address_geo_lat, _ := p.Args["address_geo_lat"].(string)
				address_geo_lng, _ := p.Args["address_geo_lng"].(string)
				phone, _ := p.Args["phone"].(string)
				website, _ := p.Args["website"].(string)
				company_name, _ := p.Args["company_name"].(string)
				company_catch_phrase, _ := p.Args["company_catch_phrase"].(string)
				company_bs, _ := p.Args["company_bs"].(string)
				row, err := db.Exec("UPDATE users SET name = $1, username= $2, email= $3, address_street= $4,address_suite= $5, address_city=$6, address_zipcode=$7, address_geo_lat=$8, address_geo_lng= $9, phone= $10, website= $11, company_name= $12, company_catch_phrase= $13, company_bs=$14 WHERE id = $15 RETURNING id", name, username, email, address_street, address_suite, address_city, address_zipcode, address_geo_lat, address_geo_lng, phone, website, company_name, company_catch_phrase, company_bs, id)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				rowsAffected, _ := row.RowsAffected()
				if rowsAffected == 0 {
					return nil, errors.New("user not found")
				}
				return row, nil
			},
		},
		"updatePost": &graphql.Field{
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
