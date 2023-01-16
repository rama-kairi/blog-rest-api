package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogStore struct {
	Blogs []Blog
}

// Get all blogs
func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Get all blogs")
}

// Get a blog
func GetBlog(w http.ResponseWriter, r *http.Request) {
	// Get blog id from url
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Blog id: %d", id)
}

// Create a blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var blog Blog
	// Decode the request body into the struct.
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Create blog")
	log.Printf("Blog: %#v ", blog)
	fmt.Fprintf(w, blog.Body)
}

func main() {
	http.HandleFunc("/blogs", GetAllBlogs)
	http.HandleFunc("/blog/", GetBlog)
	http.HandleFunc("/blog/create", CreateBlog)

	// Hello {name} REST API
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Query())
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Hello, %s Your age is %s!", name, age)
	})

	log.Println("Listening on :8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Folder Structure of  REST API
// -> Services
// -> Controllers
// -> Routes
// -> Models

// Url Query Parameters -> http://localhost:8080/hello?name=John&age=20&gendar=male
