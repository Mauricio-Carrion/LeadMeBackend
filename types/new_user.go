package types

import "github.com/google/uuid"

type NewUser struct {
	Name string
	Email string
	Phone string
	Password string
	Active bool
	IsAdmin bool
	Companie_uuid uuid.UUID
}