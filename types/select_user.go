package types

import "github.com/google/uuid"

type SelectUser struct {
	Uuid string
	Name string
	Email string
	Active bool
	IsAdmin bool
	Companie_uuid uuid.UUID
}