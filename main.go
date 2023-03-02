package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/syywu/go-graphql/db"
	"github.com/syywu/go-graphql/mutation"
	"github.com/syywu/go-graphql/query"

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

func main() {

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
	query := `
	{
		posts{
			id
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
