package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config := NewConfig()

	imageSourceUrl, err := url.Parse("https://"+ config.blobAccount  + ".blob.core.windows.net/" +
		config.blobContainer  + "/")
	if err != nil {
		log.Fatalf("Invalid image source URL | %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", home)
	r.Get("/home", home)

	r.Get("/animal", animal)

	fmt.Println(imageSourceUrl)

	serverErr := http.ListenAndServe(":"+strconv.Itoa(5555), r)

	if serverErr != nil {
		log.Fatalf("server failed: %v", serverErr)
	}
}

// home home page.
func home(w http.ResponseWriter, r *http.Request) {

}

// animal animal page.
func animal(w http.ResponseWriter, r *http.Request) {

}