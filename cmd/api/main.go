package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/camphotos/identity/pkg/handlers"
	"github.com/camphotos/identity/pkg/repository"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello World")

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", "postgres://identity:identity@database:5432/identity?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	http.HandleFunc("GET /v1/users", func(w http.ResponseWriter, r *http.Request) {

		if email := r.URL.Query().Get("email"); email != "" {
			handlers.GetUserByEmailHandler(w, r, userRepo)
		} else {
			handlers.GetUsers(w, r, userRepo)
		}
	})
	http.HandleFunc("GET /v1/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r, userRepo)
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
