package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rama-kairi/blog-rest-api/utils"
)

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogStore struct {
	Blogs []Blog
}

func NewBlogStore() *BlogStore {
	return &BlogStore{
		Blogs: []Blog{},
	}
}

// Get all blogs
func (t BlogStore) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is GET
	if !utils.CheckMethod(r.Method, utils.GET) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Marshal t.Blogs into json
	blogs, err := json.Marshal(t.Blogs)
	if err != nil {
		log.Fatal(err)
	}

	// Write the json to the response
	utils.SuccessResponse(w, http.StatusOK, blogs)
}

// Get a blog
func (t BlogStore) GetBlog(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is GET
	if !utils.CheckMethod(r.Method, utils.GET) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get blog id from url
	id, err := utils.GetUrlParamId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid blog id")
		return
	}

	for _, blog := range t.Blogs {
		if blog.Id == id {
			blogJson, err := json.Marshal(blog)
			if err != nil {
				log.Fatal(err)
			}
			utils.SuccessResponse(w, http.StatusOK, blogJson)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Blog not found")
}

// Create a blog
func (t BlogStore) CreateBlog(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is Post
	if !utils.CheckMethod(r.Method, utils.POST) {
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
	blog.Id = t.newTodoId()

	// Append the blog to the slice
	t.Blogs = append(t.Blogs, blog)
	t.saveToJson()

	// Marshal the blog into json
	data, err := json.Marshal(blog)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal the blog into json
	utils.SuccessResponse(w, http.StatusCreated, data)
}

// Delete a blog
func (t BlogStore) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is Delete
	if !utils.CheckMethod(r.Method, utils.DELETE) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get blog id from url
	id, err := utils.GetUrlParamId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, blog := range t.Blogs {
		if blog.Id == id {
			t.Blogs = append(t.Blogs[:i], t.Blogs[i+1:]...)
			t.saveToJson()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	// Save the blogs to the json file
	t.saveToJson()

	// If the blog is not found, return 404
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Blog not found")
}
