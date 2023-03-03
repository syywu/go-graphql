package models

import "github.com/graphql-go/graphql"

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	Catchphrase string `json:"catchphrase"`
	Bs          string `json:"bs"`
}

var CompanyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"name":        &graphql.Field{Type: graphql.String},
			"catchphrase": &graphql.Field{Type: graphql.String},
			"bs":          &graphql.Field{Type: graphql.String},
		},
	},
)

var GeoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Geo",
		Fields: graphql.Fields{
			"lat": &graphql.Field{Type: graphql.String},
			"lng": &graphql.Field{Type: graphql.String},
		},
	},
)

var AddressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"street":  &graphql.Field{Type: graphql.String},
			"suite":   &graphql.Field{Type: graphql.String},
			"city":    &graphql.Field{Type: graphql.String},
			"zipcode": &graphql.Field{Type: graphql.String},
			"geo":     &graphql.Field{Type: GeoType},
		},
	},
)
