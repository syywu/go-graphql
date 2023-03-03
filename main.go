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
