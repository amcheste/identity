package handlers

import (
	"encoding/json"
	"github.com/camphotos/identity/pkg/repository"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request, repo *repository.UserRepository) {

	users, err := repo.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		log.Printf("Error fetching users: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
