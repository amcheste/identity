package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/camphotos/identity/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetAllUsers_Success(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	// Expected users
	expectedUsers := []models.User{
		{
			ID:           uuid.New(),
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "johndoe@example.com",
			Status:       models.Active,
			TimeCreated:  "2023-01-01 12:00:00",
			TimeModified: "2023-01-01 12:00:00",
		},
	}

	// Mock query results
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "status", "time_created", "time_modified"}).
		AddRow(expectedUsers[0].ID.String(), expectedUsers[0].FirstName, expectedUsers[0].LastName, expectedUsers[0].Email, expectedUsers[0].Status, expectedUsers[0].TimeCreated, expectedUsers[0].TimeModified)
	mock.ExpectQuery("SELECT id, first_name, last_name, email, status, time_created, time_modified FROM users").WillReturnRows(rows)

	// Create repository
	repo := NewUserRepository(db)

	// Call GetAllUsers
	users, err := repo.GetAllUsers()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetAllUsers_Error(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	// Simulate query error
	mock.ExpectQuery("SELECT id, first_name, last_name, email, status, time_created, time_modified FROM users").WillReturnError(sql.ErrConnDone)

	// Create repository
	repo := NewUserRepository(db)

	// Call GetAllUsers
	users, err := repo.GetAllUsers()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, users)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetAllUsers_ScanError(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	// Simulate a scan error by creating rows with an incorrect number of columns
	rows := sqlmock.NewRows([]string{"id", "first_name"}). // Missing other columns
								AddRow("invalid-uuid", "John")
	mock.ExpectQuery("SELECT id, first_name, last_name, email, status, time_created, time_modified FROM users").
		WillReturnRows(rows)

	// Create repository
	repo := NewUserRepository(db)

	// Call GetAllUsers
	users, err := repo.GetAllUsers()

	// Assertions
	assert.Error(t, err, "expected scan error but got none")
	assert.Nil(t, users, "expected users to be nil on scan error")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Test GetUserById with a valid user ID
func TestGetUserById_Success(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Expected user data
	userID := uuid.New()
	expectedUser := models.User{
		ID:           userID,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		Status:       "active",
		TimeCreated:  "2023-01-01 12:00:00",
		TimeModified: "2023-01-01 12:00:00",
	}

	// Convert UUID to string for DB mock
	userIDStr := userID.String()

	// Set up the mock expectation
	query := `
		SELECT id, first_name, last_name, email, status, time_created, time_modified 
		FROM users 
		WHERE id = \$1
	`
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "status", "time_created", "time_modified"}).
		AddRow(userIDStr, expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.Status, expectedUser.TimeCreated, expectedUser.TimeModified)

	mock.ExpectQuery(query).WithArgs(userIDStr).WillReturnRows(rows)

	// Create the repository
	repo := &UserRepositoryImpl{db: db}

	// Call the function
	user, err := repo.GetUserById(userIDStr)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.FirstName, user.FirstName)
	assert.Equal(t, expectedUser.LastName, user.LastName)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Status, user.Status)

	// Ensure expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Test GetUserById when the user is not found
func TestGetUserById_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := `
		SELECT id, first_name, last_name, email, status, time_created, time_modified 
		FROM users 
		WHERE id = \$1
	`

	mock.ExpectQuery(query).WithArgs("non-existent-id").WillReturnError(sql.ErrNoRows)

	repo := &UserRepositoryImpl{db: db}

	user, err := repo.GetUserById("non-existent-id")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Test GetUserById with a database error
func TestGetUserById_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	query := `
		SELECT id, first_name, last_name, email, status, time_created, time_modified 
		FROM users 
		WHERE id = \$1
	`

	mock.ExpectQuery(query).WithArgs("some-id").WillReturnError(errors.New("database error"))

	repo := &UserRepositoryImpl{db: db}

	user, err := repo.GetUserById("some-id")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "database error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}
