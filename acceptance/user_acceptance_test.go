package acceptance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/camphotos/identity/pkg/models"
)

func TestGetAllUsers(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/v1/users")
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected response code: %d %s", resp.StatusCode, resp.Status)
	}

	// Step 1: Read the response body
	body, err := io.ReadAll(resp.Body) // Replaced ioutil.ReadAll with io.ReadAll for Go 1.16+
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Step 2: Print the raw JSON response
	fmt.Printf("Raw JSON Response:\n%s\n", body)

	var users []models.User
	if err := json.Unmarshal(body, &users); err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// User 1
	if users[0].FirstName != "John" {
		t.Errorf("Expected John as the first name, got %s ", users[0].FirstName)
	}
	if users[0].LastName != "Doe" {
		t.Errorf("Expected Doe as the last name, got %s ", users[0].LastName)
	}
	if users[0].Email != "johndoe@mail.com" {
		t.Errorf("Expected johndoe@mail.com as the email, got %s ", users[0].Email)
	}
	if users[0].Status != "ACTIVE" {
		t.Errorf("Expected ACTIVE as the status, got %s ", users[0].Status)
	}

	// User 2
	if users[1].FirstName != "Jane" {
		t.Errorf("Expected Jane as the first name, got %s ", users[1].FirstName)
	}
	if users[1].LastName != "Doe" {
		t.Errorf("Expected Doe as the last name, got %s ", users[1].LastName)
	}
	if users[1].Email != "janedoe@mail.com" {
		t.Errorf("Expected janedoe@mail.com as the email, got %s ", users[1].Email)
	}
	if users[1].Status != "ACTIVE" {
		t.Errorf("Expected ACTIVE as the status, got %s ", users[1].Status)
	}

}

func TestGetUserByID(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/v1/users/1551e20d-d4af-47ba-89ba-9a53830f0852")
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected response code: %d %s", resp.StatusCode, resp.Status)
	}

	// Step 1: Read the response body
	body, err := io.ReadAll(resp.Body) // Replaced ioutil.ReadAll with io.ReadAll for Go 1.16+
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Step 2: Print the raw JSON response
	fmt.Printf("Raw JSON Response:\n%s\n", body)

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// User 1
	if user.FirstName != "John" {
		t.Errorf("Expected John as the first name, got %s ", user.FirstName)
	}
	if user.LastName != "Doe" {
		t.Errorf("Expected Doe as the last name, got %s ", user.LastName)
	}
	if user.Email != "johndoe@mail.com" {
		t.Errorf("Expected johndoe@mail.com as the email, got %s ", user.Email)
	}
	if user.Status != "ACTIVE" {
		t.Errorf("Expected ACTIVE as the status, got %s ", user.Status)
	}

}
