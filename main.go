package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("GET /users", getUsers)
	http.HandleFunc("GET /users/{id}", getUser)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	// Convert idString to integer (handle errors appropriately)
	// ...
	fmt.Println(idString)

	user := User{ID: 1, Name: "Alice"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
