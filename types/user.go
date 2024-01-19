package types

import "github.com/google/uuid"

type User struct {
	Uuid string
	Name string
	Email string
	Password string
	Active bool
	IsAdmin bool
	Jid string
	Companie_uuid uuid.UUID
}