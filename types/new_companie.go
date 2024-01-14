package types

import "github.com/google/uuid"

type NewCompanie struct {
	Name string
	Razao string
	Cpf_cnpj string
	Active bool
	Companie_uuid uuid.UUID
}