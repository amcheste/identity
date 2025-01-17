package handlers

import (
	"encoding/json"
	"github.com/camphotos/identity/pkg/repository"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request, repo repository.UserRepository) {

	users, err := repo.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		log.Printf("Error fetching users: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, repo repository.UserRepository) {
	id := r.PathValue("id")
	log.Printf(id)
	user, err := repo.GetUserById(id)
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError) //TODO: Add id
		log.Printf("Error fetching user: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func GetUserByEmailHandler(w http.ResponseWriter, r *http.Request, repo repository.UserRepository) {
	// Extract email from the query parameter
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email query parameter is required", http.StatusBadRequest)
		return
	}

	// Retrieve the user by email
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Failed to fetch user by email", http.StatusNotFound)
		log.Printf("Error fetching user by email: %v", err)
		return
	}

	// Write the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
