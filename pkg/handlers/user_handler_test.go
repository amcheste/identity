package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/camphotos/identity/pkg/models"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockUserRepository implements the UserRepository interface for testing
type MockUserRepository struct {
	GetAllUsersFn func() ([]models.User, error)
}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	return m.GetAllUsersFn()
}

func TestGetUsers_Success(t *testing.T) {
	// Mock successful response
	mockRepo := &MockUserRepository{
		GetAllUsersFn: func() ([]models.User, error) {
			return []models.User{
				{
					ID:        uuid.New(),
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@example.com",
					Status:    models.Active,
				},
			}, nil
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	GetUsers(rec, req, mockRepo)

	// Assertions
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	var users []models.User
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(users) != 1 || users[0].FirstName != "John" {
		t.Errorf("Unexpected users data: %+v", users)
	}
}

func TestGetUsers_Failure(t *testing.T) {
	// Mock error response
	mockRepo := &MockUserRepository{
		GetAllUsersFn: func() ([]models.User, error) {
			return nil, errors.New("database error")
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	GetUsers(rec, req, mockRepo)

	// Assertions
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, res.StatusCode)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(res.Body)

	expectedError := "Failed to fetch users"
	if !bytes.Contains(body.Bytes(), []byte(expectedError)) {
		t.Errorf("Expected error message %q in response, got %q", expectedError, body.String())
	}
}
