package models

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

// Test the `String()` method of the `Status` type.
func TestStatusString(t *testing.T) {
	// Test for the ACTIVE status
	statusActive := Active
	expected := "ACTIVE"
	actual := statusActive.String()
	if expected != actual {
		t.Errorf("Expected ACTIVE to be 'ACTIVE', got '%s'", actual)
	}

	// Test for the INACTIVE status
	statusInactive := Inactive
	expected = "IN_ACTIVE"
	actual = statusInactive.String()
	if expected != actual {
		t.Errorf("Expected INACTIVE to be 'IN_ACTIVE', got '%s'", actual)
	}
}

// Test the User struct initialization and its behavior.
func TestUser(t *testing.T) {
	// Creating a new UUID for the user
	testUUID := uuid.New()

	// Creating a timestamp string as the time when user is created (using the current time)
	currentTime := time.Now().Format(time.RFC3339)

	// Initializing a User object
	user := User{
		ID:           testUUID,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "johndoe@example.com",
		Status:       Active,
		TimeCreated:  currentTime,
		TimeModified: currentTime,
	}

	// Checking that the user struct is initialized as expected.
	if user.ID != testUUID {
		t.Errorf("Expected the user's UUID to match the given UUID, got '%v'", user.ID)
	}
	if user.FirstName != "John" {
		t.Errorf("Expected the first name to be 'John', got '%s'", user.FirstName)
	}
	if user.LastName != "Doe" {
		t.Errorf("Expected the last name to be 'Doe', got '%s'", user.LastName)
	}
	if user.Email != "johndoe@example.com" {
		t.Errorf("Expected the email to be 'johndoe@example.com', got '%s'", user.Email)
	}
	if user.Status != Active {
		t.Errorf("Expected the user status to be 'ACTIVE', got '%s'", user.Status)
	}
	if user.TimeCreated != currentTime {
		t.Errorf("Expected the creation time to match the given time, got '%s'", user.TimeCreated)
	}
	if user.TimeModified != currentTime {
		t.Errorf("Expected the modification time to match the given time, got '%s'", user.TimeModified)
	}
}

// Test the default values for an empty User struct
func TestEmptyUser(t *testing.T) {
	// Creating an empty user struct
	user := User{}

	// Check that all default values are set as expected (for zero values)
	if user.ID != uuid.Nil {
		t.Errorf("Expected user ID to be empty (nil UUID) by default, got '%v'", user.ID)
	}
	if user.FirstName != "" {
		t.Errorf("Expected user first name to be empty by default, got '%s'", user.FirstName)
	}
	if user.LastName != "" {
		t.Errorf("Expected user last name to be empty by default, got '%s'", user.LastName)
	}
	if user.Email != "" {
		t.Errorf("Expected user email to be empty by default, got '%s'", user.Email)
	}
	if user.Status != "" {
		t.Errorf("Expected user status to be empty by default, got '%s'", user.Status)
	}
	if user.TimeCreated != "" {
		t.Errorf("Expected user time_created to be empty by default, got '%s'", user.TimeCreated)
	}
	if user.TimeModified != "" {
		t.Errorf("Expected user time_modified to be empty by default, got '%s'", user.TimeModified)
	}
}
