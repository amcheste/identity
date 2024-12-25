package integration

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/camphotos/identity/pkg/repository"
	"testing"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupTestContainer(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	ctx := context.Background()

	// Define PostgreSQL container request
	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine3.20", // Use an official Postgres image
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "identity",
			"POSTGRES_PASSWORD": "identity",
			"POSTGRES_DB":       "identity",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	// Start the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}

	// Get the container's mapped port
	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get container port: %v", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	// Connect to the database
	dsn := "postgres://identity:identity@%s:%s/identity?sslmode=disable"
	db, err := sql.Open("postgres", fmt.Sprintf(dsn, host, mappedPort.Port()))
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Return a cleanup function
	cleanup := func() {
		db.Close()
		container.Terminate(ctx)
	}

	return db, cleanup
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	// Setup container and clean up afterward
	db, cleanup := setupTestContainer(t)
	defer cleanup()

	// Run database setup
	setupDatabase(t, db)

	// Insert test data
	if _, err := db.Exec(`INSERT INTO users (id, first_name, last_name, email) VALUES 
		('11111111-1111-1111-1111-111111111111', 'John', 'Doe', 'johndoe@mail.com'),
		('22222222-2222-2222-2222-222222222222', 'Jane', 'Doe', 'janedoe@mail.com');`); err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Initialize repository
	repo := repository.NewUserRepository(db)

	// Test GetAllUsers
	users, err := repo.GetAllUsers()
	if err != nil {
		t.Fatalf("Error fetching users: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	expected := map[string]string{
		"11111111-1111-1111-1111-111111111111": "John Doe",
		"22222222-2222-2222-2222-222222222222": "Jane Doe"}

	for _, user := range users {
		if name := user.FirstName + " " + user.LastName; expected[user.ID.String()] != name {
			t.Errorf("Unexpected user: got %s, want %s", name, expected[user.ID.String()])
		}
	}

}

func setupDatabase(t *testing.T, db *sql.DB) {
	t.Helper()

	schema := `
		CREATE TYPE status AS ENUM ('ACTIVE', 'IN_ACTIVE');
		CREATE TABLE users (
    		id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    		first_name VARCHAR(100),
    		last_name VARCHAR(100),
    		email VARCHAR(150),
    		status status NOT NULL DEFAULT 'ACTIVE',
    		time_created timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		time_modified timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX users_email_index ON users(email);
	`
	if _, err := db.Exec(schema); err != nil {
		t.Fatalf("Failed to set up database schema: %v", err)
	}
}
