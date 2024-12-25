package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/camphotos/identity/pkg/models"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
}

// UserRepositoryImpl is a concrete implementation of UserRepository using a PostgreSQL database
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository implementation
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetAllUsers retrieves all users from the database
func (repo *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	rows, err := repo.db.Query("SELECT id, first_name, last_name, email, status, time_created, time_modified FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var id string
		if err := rows.Scan(&id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.TimeCreated, &user.TimeModified); err != nil {
			return nil, err
		}
		user.ID, _ = uuid.Parse(id)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
