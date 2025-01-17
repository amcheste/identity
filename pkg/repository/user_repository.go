package repository

import (
	"database/sql"

	"github.com/camphotos/identity/pkg/models"
	"github.com/google/uuid"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
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

// GetUserByID retrieves a single user by their ID
func (repo *UserRepositoryImpl) GetUserById(id string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, email, status, time_created, time_modified 
		FROM users 
		WHERE id = $1
	`

	var user models.User
	var userID string

	err := repo.db.QueryRow(query, id).Scan(
		&userID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.TimeCreated,
		&user.TimeModified,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	user.ID, err = uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, email, status, time_created, time_modified 
		FROM users 
		WHERE email = $1
	`

	var user models.User
	var id string

	err := repo.db.QueryRow(query, email).Scan(
		&id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.TimeCreated,
		&user.TimeModified,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	user.ID, err = uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
