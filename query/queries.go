package query

import (
	"fmt"

	"github.com/graphql-go/graphql"
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
				"id": &graphql.ArgumentConfig{Type: graphql.ID},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return models.GetPostByID(p)
			},
		},
		"posts": &graphql.Field{
			Type:        graphql.NewList(models.PostType),
			Description: "Get All Posts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return models.GetPosts(p)
			},
		},
		"user": &graphql.Field{
			Type:        models.UserType,
			Description: "Get User by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Print(p.Source.(models.User))
				return models.GetUserByID(p)
			},
		},
		"users": &graphql.Field{
			Type:        graphql.NewList(models.UserType),
			Description: "Get All Users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return models.GetUsers(p)
			},
		},
	},
})
