package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/syywu/go-graphql/db"
	"github.com/syywu/go-graphql/models"
	"github.com/syywu/go-graphql/mutation"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

// users
/*
  {
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
    "address": {
      "street": "Kulas Light",
      "suite": "Apt. 556",
      "city": "Gwenborough",
      "zipcode": "92998-3874",
      "geo": {
        "lat": "-37.3159",
        "lng": "81.1496"
      }
    },
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org",
    "company": {
      "name": "Romaguera-Crona",
      "catchPhrase": "Multi-layered client-server neural-net",
      "bs": "harness real-time e-markets"
    }
  },
*/

// define our schema- define what fields to return to us for when making queries
var queryType = graphql.NewObject(graphql.ObjectConfig{
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
				posts := []models.Post{}

				for rows.Next() {
					var post models.Post
					err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
					if err != nil {
						return nil, err
					}
					posts = append(posts, post)
				}
				defer rows.Close()
				return posts, nil
			},
		},
	},
})

func main() {

	db.CreatePostsTable()
	// CreateCommentsTable()

	router := chi.NewRouter()

	// rootQuery- where to start
	// defines a schema config
	schemaConfig := graphql.SchemaConfig{Query: queryType, Mutation: mutation.MutationType}
	// creates schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("failed to create new graphql schema", err)
	}

	// router.Handle("/graphql", &GraphQLHandler{Schema: schema})

	// query- to get back our query
	query := `
	{
		posts{
			id
			title
		}
	}
	`
	// create a params struct which contains a reference to our defined Schema as well as our RequestString request.
	params := graphql.Params{Schema: schema, RequestString: query}
	//  execute the request and the results of the request are populated into r
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatal("failed to execute graphql operations", r.Errors)
	}
	//  Marshal the response into JSON and print it out to our console
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
