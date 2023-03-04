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

// func populate() []models.User {
// 	geo := &models.Geo{Lat: "37484", Lng: "369864"}
// 	address := &models.Address{Street: "Rotherhithe Street", Suite: "200", City: "London", Zipcode: "SE6 1NZ", Geo: *geo}
// 	company := &models.Company{Name: "ahbakjf", Catchphrase: "hdsu", Bs: "jbfibgf"}
// 	user := models.User{Name: "John", Username: "jdh23", Email: "jgfgfusdc", Address: *address, Phone: "637971333", Website: "jahfdkgbuafdua.com", Company: *company}

// 	geo2 := &models.Geo{Lat: "2222", Lng: "37372"}
// 	address2 := &models.Address{Street: "cdslgofuagod", Suite: "200", City: "Swansea", Zipcode: "SE74 1MZ", Geo: *geo2}
// 	company2 := &models.Company{Name: "Atos", Catchphrase: "djaffaol", Bs: "dhkcfugakd"}
// 	user2 := models.User{Name: "jfhwauf", Username: "gd76", Email: "ajdda@jshc.com", Address: *address2, Phone: "726924222", Website: "jhfgas.com", Company: *company2}
// 	var users []models.User
// 	users = append(users, user, user2)

// 	return users
// }

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
}
