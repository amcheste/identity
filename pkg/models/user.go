package models

import (
	"github.com/google/uuid"
)

type Status string

const (
	Active   Status = "ACTIVE"
	Inactive Status = "IN_ACTIVE"
)

func (c Status) String() string {
	return string(c)
}

type User struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Status       Status    `json:"status"`
	TimeCreated  string    `json:"time_created"`
	TimeModified string    `json:"time_modified"`
}
