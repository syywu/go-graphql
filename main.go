package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/syywu/go-graphql/db"
	"github.com/syywu/go-graphql/mutation"
	"github.com/syywu/go-graphql/query"

	_ "github.com/lib/pq"
)

func main() {

	db.CreateUsersTable()
	db.CreatePostsTable()

	r := chi.NewRouter()

	// rootQuery- where to start
	// creates schema and defines a schema config
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.QueryType,
		Mutation: mutation.MutationType,
	})
	if err != nil {
		log.Fatal("failed to create new graphql schema", err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.Handle("/", graphQLHandler)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
	// query- to get back our query
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "Root", Fields: fields}),
	})

	if err != nil {
		log.Fatal(err)
	}

	query := `
	{
		users{
			name
			address{
				geo{
					lat
					lng
				}
			}
			company{
				name
				bs
			}
		}
	}
	`
	// create a params struct which contains a reference to our defined Schema as well as our RequestString request.
	params := graphql.Params{Schema: schema, RequestString: query}

	req := graphql.Do(params)
	if len(req.Errors) > 0 {
		log.Fatal("failed to execute graphql operations", req.Errors)
	}

	rJSON, _ := json.Marshal(req)
	fmt.Printf("%s \n", rJSON)
*/
