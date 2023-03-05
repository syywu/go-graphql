package models

import (
	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
)

func GetPosts(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	defer rows.Close()
	posts := []Post{}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetUsers(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Address.Street, &user.Address.Suite, &user.Address.City, &user.Address.Zipcode, &user.Address.Geo.Lat, &user.Address.Geo.Lng, &user.Phone, &user.Website, &user.Company.Name, &user.Company.Catchphrase, &user.Company.Bs)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
