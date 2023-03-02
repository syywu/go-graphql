package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/graphql?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreatePostsTable() {
	db := OpenConnection()
	const createPostTable = `
	CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		userId INT NOT NULL,
		title TEXT NOT NULL,
		body TEXT
	);
	`
	_, err := db.Exec(createPostTable)
	if err != nil {
		log.Fatal("cannot create posts table", err)
	}
	defer db.Close()

}

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var postType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.Int},
			"userId": &graphql.Field{Type: graphql.Int},
			"title":  &graphql.Field{Type: graphql.String},
			"body":   &graphql.Field{Type: graphql.String},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        postType,
			Description: "Create a New Post",
			Args: graphql.FieldConfigArgument{
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := OpenConnection()
				userid, _ := p.Args["userid"].(int)
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				var post Post
				_, err := db.Exec("INSERT INTO posts (userid, title, body) VALUES ($1, $2, $3) RETURNING id", userid, title, body)
				if err != nil {
					return nil, err
				}
				defer db.Close()
				return post, nil
			},
		},
		"update": &graphql.Field{
			Type:        postType,
			Description: "Update a Post",
			Args: graphql.FieldConfigArgument{
				"id":     &graphql.ArgumentConfig{Type: graphql.Int},
				"userid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"title":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"body":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := OpenConnection()
				id, _ := p.Args["id"].(int)
				userid, _ := p.Args["userid"].(int)
				title, _ := p.Args["title"].(string)
				body, _ := p.Args["body"].(string)
				row, err := db.Exec("UPDATE posts SET userid = $1, title = $2, body = $3 WHERE id = $4", userid, title, body, id)
				if err != nil {
					return nil, err
				}
				row.RowsAffected()
				defer db.Close()
				return row, nil
			},
		},
	},
})

// define our schema- define what fields to return to us for when making queries
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type:        postType,
			Description: "Get Post by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := OpenConnection()
				id, _ := p.Args["id"].(int)
				row := db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
				defer db.Close()

				var post Post
				err := row.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
				if err != nil {
					return nil, err
				}
				return post, nil
			},
		},
		"posts": &graphql.Field{
			Type:        graphql.NewList(postType),
			Description: "Get All Posts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db := OpenConnection()
				rows, err := db.Query("SELECT * FROM posts")
				if err != nil {
					return nil, err
				}
				defer db.Close()
				posts := []Post{}

				for rows.Next() {
					var post Post
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

	CreatePostsTable()
	// CreateCommentsTable()

	router := chi.NewRouter()

	// rootQuery- where to start
	// defines a schema config
	schemaConfig := graphql.SchemaConfig{Query: queryType, Mutation: mutationType}
	// creates schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("failed to create new graphql schema", err)
	}

	// router.Handle("/graphql", &GraphQLHandler{Schema: schema})

	// query- to get back our query
	query := `
	{
  		post(id:1){
			userId
			title
			body
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
