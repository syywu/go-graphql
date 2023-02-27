package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
