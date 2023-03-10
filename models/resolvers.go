package models

import (
	"errors"

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

func CreatePost(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	title, _ := p.Args["title"].(string)
	body, _ := p.Args["body"].(string)
	userid, _ := p.Args["userid"].(string)
	var post Post
	_, err := db.Exec("INSERT INTO posts (title, body, userid) VALUES ($1, $2, $3) RETURNING id", title, body, userid)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return post, nil
}

func UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	id, _ := p.Args["id"].(string)
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
	row, err := db.Exec("UPDATE users SET name = $1, username= $2, email= $3, address_street= $4,address_suite= $5, address_city=$6, address_zipcode=$7, address_geo_lat=$8, address_geo_lng= $9, phone= $10, website= $11, company_name= $12, company_catch_phrase= $13, company_bs=$14 WHERE id = $15 RETURNING id", name, username, email, address_street, address_suite, address_city, address_zipcode, address_geo_lat, address_geo_lng, phone, website, company_name, company_catch_phrase, company_bs, id)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return row, nil
}

func UpdatePost(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	id, _ := p.Args["id"].(string)
	userid, _ := p.Args["userid"].(string)
	title, _ := p.Args["title"].(string)
	body, _ := p.Args["body"].(string)
	row, err := db.Exec("UPDATE posts SET title = $1, body = $2, userid = $3 WHERE id = $4 RETURNING *", title, body, userid, id)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("post not found")
	}
	return row, nil
}

func DeleteUser(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	id, _ := p.Args["id"].(string)
	row, err := db.Exec("DELETE FROM users WHERE id =$1 RETURNING id", id)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return true, nil
}

func DeletetPost(p graphql.ResolveParams) (interface{}, error) {
	db := db.OpenConnection()
	id, _ := p.Args["id"].(string)
	row, err := db.Exec("DELETE FROM posts WHERE id =$1 RETURNING id", id)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("post not found")
	}
	return true, nil
}
