package main

import (
	"log"
	"net/http"

	"github.com/rama-kairi/blog-rest-api/controllers"
)

func main() {
	mux := http.NewServeMux()
	newBlog := controllers.NewBlogStore()

	mux.HandleFunc("/blog/all/", newBlog.GetAllBlogs)
	mux.HandleFunc("/blog/one/", newBlog.GetBlog)
	mux.HandleFunc("/blog/create/", newBlog.CreateBlog)
	mux.HandleFunc("/blog/delete/", newBlog.DeleteBlog)

	log.Println("Listening on :8080...")

	http.ListenAndServe(":8080", mux)
}

// Folder Structure of  REST API
// -> Services
// -> Controllers
// -> Routes
// -> Models

// Url Query Parameters -> http://localhost:8080/hello?name=John&age=20&gendar=male
