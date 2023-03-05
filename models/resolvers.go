package models

import (
	"github.com/graphql-go/graphql"
	"github.com/syywu/go-graphql/db"
)

func GetPostByID(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	id, _ := p.Args["id"].(string)
	row := db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
	defer db.Close()

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.UserId)
	if err != nil {
		return nil, err
	}
	return post, nil
}

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

func GetUserByID(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	db := db.OpenConnection()
	user := User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Address.Street, &user.Address.Suite, &user.Address.City, &user.Address.Zipcode, &user.Address.Geo.Lat, &user.Address.Geo.Lng, &user.Phone, &user.Website, &user.Company.Name, &user.Company.Catchphrase, &user.Company.Bs)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return user, nil
}
