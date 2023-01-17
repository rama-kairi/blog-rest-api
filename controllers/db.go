package controllers

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveToJson - This function saves the todo list to a json file
func (t BlogStore) saveToJson() {
	// Marshal the data from t.Blogs
	data, err := json.Marshal(t.Blogs)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("db.json", data, 0o644); err != nil {
		panic(err)
	}
}

// LoadFromJson - This function loads the todo list from a json file
func (t *BlogStore) loadFromJson() {
	// Create a file if it doesn't exist
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		_, err := os.Create("db.json")
		if err != nil {
			panic(err)
		}
	}

	data, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		// Unmarshal the data to t.Blogs
		err = json.Unmarshal(data, &t.Blogs)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No todo's found")
	}
}

// newTodoId - This function returns a new todo id
func (t BlogStore) newTodoId() int {
	if len(t.Blogs) == 0 {
		return 1
	}
	return t.Blogs[len(t.Blogs)-1].Id + 1
}
