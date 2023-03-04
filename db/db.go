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
		title TEXT NOT NULL,
		body TEXT,
		userid INT NOT NULL
	);
	`
	_, err := db.Exec(createPostTable)
	if err != nil {
		log.Fatal("cannot create posts table", err)
	}
	defer db.Close()

}

func CreateUsersTable() {
	db := OpenConnection()
	const createUsersTable = `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
    	name VARCHAR(255) NOT NULL,
    	username VARCHAR(255) NOT NULL,
    	email VARCHAR(255) NOT NULL,
   		address_street VARCHAR(255) NOT NULL,
   		address_suite VARCHAR(255) NOT NULL,
    	address_city VARCHAR(255) NOT NULL,
    	address_zipcode VARCHAR(255) NOT NULL,
    	address_geo_lat VARCHAR(255) NOT NULL,
    	address_geo_lng VARCHAR(255) NOT NULL,
    	phone VARCHAR(255) NOT NULL,
    	website VARCHAR(255) NOT NULL,
   		company_name VARCHAR(255) NOT NULL,
    	company_catch_phrase VARCHAR(255) NOT NULL,
    	company_bs VARCHAR(255) NOT NULL
	)
	`
	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatal("cannot create users table", err)
	}
	defer db.Close()
}
