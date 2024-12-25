package repository

import (
	"database/sql"
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
