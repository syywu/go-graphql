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

// comments
/*
  {
    "postId": 1,
    "id": 1,
    "name": "id labore ex et quam laborum",
    "email": "Eliseo@gardner.biz",
    "body": "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"
  },
*/

// posts
/*
  {
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  },
*/

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

func CreateCommentsTable() {
	db := OpenConnection()
	const createCommentTable = `
		CREATE TABLE IF NOT EXISTS comments(
		id SERIAL PRIMARY KEY,
		postId INT NOT NULL REFERENCES posts(id),
		name TEXT NOT NULL,
		email VARCHAR(255) NOT NULL,
		body TEXT
	);
	`
	_, err := db.Exec(createCommentTable)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {

	CreatePostsTable()
	CreateCommentsTable()

	router := chi.NewRouter()

	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"PostId": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Email": &graphql.Field{
					Type: graphql.String,
				},
				"Body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	// defines the object config
	// where to start
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// defines a schema config
	// defines which query to allow users to use when making queries on the frontend
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	// creates schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("failed to create new graphql schema", err)
	}

	// query- to get back our query
	query := `
		{
			
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatal("failed to execute graphql operations", err)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
