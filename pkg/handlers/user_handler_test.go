package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/camphotos/identity/pkg/models"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// MockUserRepository implements the UserRepository interface for testing
type MockUserRepository struct {
	GetAllUsersFn    func() ([]models.User, error)
	GetByIDFn        func(string) (*models.User, error)
	GetUserByEmailFn func(string) (*models.User, error)
}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	return m.GetAllUsersFn()
}

func (m *MockUserRepository) GetUserById(id string) (*models.User, error) {
	return m.GetByIDFn(id)
}
func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	return m.GetUserByEmailFn(email)
}

func init_db() (m *MockUserRepository) {
	tmp, err := uuid.Parse("b48e83a5-cb34-4db5-8ef3-a4423bebfe67")
	if err != nil {
		log.Fatalf("Error parsing UUID: %v", err)
	}
	mockRepo := &MockUserRepository{
		GetAllUsersFn: func() ([]models.User, error) {
			return []models.User{
				{
					ID:        uuid.New(),
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@example.com",
					Status:    models.Active, //TODO add times
				},
			}, nil
		},
		GetByIDFn: func(id string) (*models.User, error) {
			return &models.User{
				ID:           tmp,
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "johndoe@example.com",
				TimeCreated:  time.Now().String(),
				TimeModified: time.Now().String(),
			}, nil
		},
		GetUserByEmailFn: func(email string) (*models.User, error) {
			return &models.User{
				ID:           tmp,
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "johndoe@example.com",
				TimeCreated:  time.Now().String(),
				TimeModified: time.Now().String(),
			}, nil
		},
	}

	return mockRepo
}

func TestGetUsers_Success(t *testing.T) {
	mockRepo := init_db()

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

func TestGetUserById_Success(t *testing.T) {
	mockRepo := init_db()

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	GetUser(rec, req, mockRepo)

	// Assertions
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	var user models.User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if user.FirstName != "John" {
		t.Errorf("Unexpected users data: %+v", user)
	}
}

func TestGetUserById_Failure(t *testing.T) {
	// Mock error response
	mockRepo := &MockUserRepository{
		GetByIDFn: func(id string) (*models.User, error) {
			return nil, errors.New("database error")
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	GetUser(rec, req, mockRepo)

	// Assertions
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, res.StatusCode)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(res.Body)

	expectedError := "Failed to fetch user"
	if !bytes.Contains(body.Bytes(), []byte(expectedError)) {
		t.Errorf("Expected error message %q in response, got %q", expectedError, body.String())
	}
}
