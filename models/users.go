package models

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
)

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	Catchphrase string `json:"catchphrase"`
	Bs          string `json:"bs"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.ID},
			"name":     &graphql.Field{Type: graphql.String},
			"username": &graphql.Field{Type: graphql.String},
			"email":    &graphql.Field{Type: graphql.String},
			"address":  &graphql.Field{Type: AddressType},
			"phone":    &graphql.Field{Type: graphql.String},
			"website":  &graphql.Field{Type: graphql.String},
			"company":  &graphql.Field{Type: CompanyType},
			"post": &graphql.Field{
				Type: graphql.NewList(PostType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					parent, ok := p.Source.(*User)
					fmt.Print(parent.ID)
					if ok {
						db := db.OpenConnection()
						rows, err := db.Query("SELECT * FROM posts WHERE userid = $1", parent.ID)
						if err != nil {
							return nil, err
						}
						defer db.Close()
						defer rows.Close()
						posts := []Post{}

						for rows.Next() {
							var post Post
							err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserId)
							if err != nil {
								return nil, err
							}
							posts = append(posts, post)
						}
						return posts, nil
					}
					return nil, nil
				},
			},
		},
	},
)
var AddressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"street":  &graphql.Field{Type: graphql.String},
			"suite":   &graphql.Field{Type: graphql.String},
			"city":    &graphql.Field{Type: graphql.String},
			"zipcode": &graphql.Field{Type: graphql.String},
			"geo":     &graphql.Field{Type: GeoType},
		},
	},
)

var GeoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Geo",
		Fields: graphql.Fields{
			"lat": &graphql.Field{Type: graphql.String},
			"lng": &graphql.Field{Type: graphql.String},
		},
	},
)

var CompanyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"name":        &graphql.Field{Type: graphql.String},
			"catchphrase": &graphql.Field{Type: graphql.String},
			"bs":          &graphql.Field{Type: graphql.String},
		},
	},
)
