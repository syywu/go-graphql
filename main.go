package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	} else {
		log.Println("db connection established")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreatePostTable() {
	db := OpenConnection()
	var createPostTable = `
	CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		userId INT NOT NULL,
		title TEXT,
		body TEXT
	);
	`
	_, err := db.Exec(createPostTable)
	if err != nil {
		log.Fatal("cannot create posts table", err)
	}
	defer db.Close()

}

func main() {

	CreatePostTable()

	r := chi.NewRouter()

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
