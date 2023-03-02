package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	// CreateCommentsTable()

	router := chi.NewRouter()

	// rootQuery- where to start
	// defines a schema config
	schemaConfig := graphql.SchemaConfig{Query: query.QueryType, Mutation: mutation.MutationType}
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
