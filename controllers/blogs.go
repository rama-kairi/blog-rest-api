package controllers

import (
	"encoding/json"
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
		utils.Response(w, http.StatusMethodNotAllowed, nil, "Method not allowed")
	}
	// Marshal t.Blogs into json
	blogs, err := json.Marshal(t.Blogs)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, nil, "Error marshalling blogs")
	}

	// Write the json to the response
	utils.Response(w, http.StatusOK, blogs, "Blogs found")
}

// Get a blog
func (t BlogStore) GetBlog(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is GET
	if !utils.CheckMethod(r.Method, utils.GET) {
		utils.Response(w, http.StatusMethodNotAllowed, nil, "Method not allowed")
		return
	}

	// Get blog id from url
	id, err := utils.GetUrlParamId(r)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, nil, "Error getting blog id")
		return
	}

	for _, blog := range t.Blogs {
		if blog.Id == id {
			blogJson, err := json.Marshal(blog)
			if err != nil {
				log.Fatal(err)
			}
			utils.Response(w, http.StatusOK, blogJson, "Blog found")
			return
		}
	}
	utils.Response(w, http.StatusNotFound, nil, "Blog not found")
}

// Create a blog
func (t BlogStore) CreateBlog(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is Post
	if !utils.CheckMethod(r.Method, utils.POST) {
		utils.Response(w, http.StatusMethodNotAllowed, nil, "Method not allowed")
		return
	}

	var blog Blog
	// Decode the request body into the struct.
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, nil, "Error decoding blog")
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
	utils.Response(w, http.StatusCreated, data, "Blog created successfully")
}

// Delete a blog
func (t BlogStore) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	t.loadFromJson()
	// Check if the method is Delete
	if !utils.CheckMethod(r.Method, utils.DELETE) {
		utils.Response(w, http.StatusMethodNotAllowed, nil, "Method not allowed")
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
			utils.Response(w, http.StatusNoContent, nil, "Blog deleted successfully")
			return
		}
	}

	// Save the blogs to the json file
	t.saveToJson()

	// If the blog is not found, return 404
	utils.Response(w, http.StatusNotFound, nil, "Blog not found")
}
