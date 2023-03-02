package main

import (
	"encoding/json"
	"fmt"
	"log"

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

// func OpenConnection() *sql.DB {
// 	db, err := sql.Open("postgres", "postgres://user:password@localhost/graphql?sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Cannot connect to db", err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }

// func CreatePostsTable() {
// 	db := OpenConnection()
// 	const createPostTable = `
// 	CREATE TABLE IF NOT EXISTS posts(
// 		id SERIAL PRIMARY KEY,
// 		userId INT NOT NULL,
// 		title TEXT NOT NULL,
// 		body TEXT
// 	);
// 	`
// 	_, err := db.Exec(createPostTable)
// 	if err != nil {
// 		log.Fatal("cannot create posts table", err)
// 	}
// 	defer db.Close()

// }

// func CreateCommentsTable() {
// 	db := OpenConnection()
// 	const createCommentTable = `
// 		CREATE TABLE IF NOT EXISTS comments(
// 		id SERIAL PRIMARY KEY,
// 		postId INT NOT NULL REFERENCES posts(id),
// 		name TEXT NOT NULL,
// 		email VARCHAR(255) NOT NULL,
// 		body TEXT
// 	);
// 	`
// 	_, err := db.Exec(createCommentTable)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// }

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//	type Comment struct {
//		ID     int    `json:"id"`
//		PostID int    `json:"postid"`
//		Name   string `json:"name"`
//		Email  string `json:"email"`
//		Body   string `json:"body"`
//	}

func populatePosts() []Post {
	post := Post{
		ID:     1,
		UserId: 2,
		Title:  "first",
		Body:   "hello world",
	}
	post2 := Post{
		ID:     2,
		UserId: 1,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}
	post3 := Post{
		ID:     3,
		UserId: 25,
		Title:  "rem alias distinctio quo quis",
		Body:   "ullam consequatur ut\nomnis quis sit vel consequuntur\nipsa eligendi ipsum molestiae et omnis error nostrum\nmolestiae illo tempore quia et distinctio",
	}
	var posts []Post
	posts = append(posts, post)
	posts = append(posts, post2)
	posts = append(posts, post3)
	return posts
}

// func populateComments() []Comment {
// 	comment := Comment{ID: 1, PostID: 2, Name: "Sue", Email: "sue@gmail.com", Body: "first comment"}
// 	var comments []Comment
// 	comments = append(comments, comment)
// 	return comments
// }

func main() {

	// CreatePostsTable()
	// CreateCommentsTable()

	// router := chi.NewRouter()

	allPosts := populatePosts()

	// comments := populateComments()

	// db := OpenConnection()

	// var commentType = graphql.NewObject(
	// 	graphql.ObjectConfig{
	// 		Name: "Comment",
	// 		Fields: graphql.Fields{
	// 			"ID": &graphql.Field{
	// 				Type: graphql.Int,
	// 			},
	// 			"PostId": &graphql.Field{
	// 				Type: graphql.Int,
	// 			},
	// 			"Name": &graphql.Field{
	// 				Type: graphql.String,
	// 			},
	// 			"Email": &graphql.Field{
	// 				Type: graphql.String,
	// 			},
	// 			"Body": &graphql.Field{
	// 				Type: graphql.String,
	// 			},
	// 		},
	// 	},
	// )

	var postType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Post",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"UserId": &graphql.Field{
					Type: graphql.Int,
				},
				"Title": &graphql.Field{
					Type: graphql.String,
				},
				"Body": &graphql.Field{
					Type: graphql.String,
				},
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
					"userid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"body": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post := Post{
						UserId: p.Args["userid"].(int),
						Title:  p.Args["title"].(string),
						Body:   p.Args["body"].(string),
					}
					allPosts = append(allPosts, post)
					fmt.Print(allPosts)
					return allPosts, nil
				},
			},
		},
	})

	// define our schema- define what fields to return to us for when making queries
	fields := graphql.Fields{
		"post": &graphql.Field{
			Type:        postType,
			Description: "Get Post by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, post := range allPosts {
						if int(post.ID) == id {
							return post, nil
						}
					}
				}
				return nil, nil
			},
		},
		"posts": &graphql.Field{
			Type:        graphql.NewList(postType),
			Description: "Get All Posts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return allPosts, nil
			},
		},
	}

	// defines the object config
	// rootQuery- where to start
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// defines a schema config
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: mutationType}
	// creates schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("failed to create new graphql schema", err)
	}

	// query- to get back our query
	query := `
	{
  		posts{
			Title
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

	// fmt.Println("listening on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", router))

}
