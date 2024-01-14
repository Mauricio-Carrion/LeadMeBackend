package types

import "github.com/google/uuid"

type User struct {
	Uuid uuid.UUID
	Name string
	Email string
	Password string
	Active bool
	IsAdmin bool
	Companie_uuid uuid.UUID
}