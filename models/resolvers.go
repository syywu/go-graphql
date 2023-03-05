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
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.Userid)
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
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Userid)
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

func CreateUser(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	name, _ := p.Args["name"].(string)
	username, _ := p.Args["username"].(string)
	email, _ := p.Args["email"].(string)
	address_street, _ := p.Args["address_street"].(string)
	address_suite, _ := p.Args["address_suite"].(string)
	address_city, _ := p.Args["address_city"].(string)
	address_zipcode, _ := p.Args["address_zipcode"].(string)
	address_geo_lat, _ := p.Args["address_geo_lat"].(string)
	address_geo_lng, _ := p.Args["address_geo_lng"].(string)
	phone, _ := p.Args["phone"].(string)
	website, _ := p.Args["website"].(string)
	company_name, _ := p.Args["company_name"].(string)
	company_catch_phrase, _ := p.Args["company_catch_phrase"].(string)
	company_bs, _ := p.Args["company_bs"].(string)
	var post Post
	_, err := db.Exec("INSERT INTO users (name, username, email, address_street,address_suite, address_city, address_zipcode, address_geo_lat, address_geo_lng, phone, website, company_name, company_catch_phrase, company_bs) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id", name, username, email, address_street, address_suite, address_city, address_zipcode, address_geo_lat, address_geo_lng, phone, website, company_name, company_catch_phrase, company_bs)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return post, nil
}
