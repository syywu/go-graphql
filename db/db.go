package db

import (
	"database/sql"
	"log"
)

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
