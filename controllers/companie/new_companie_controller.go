package companie

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"gorm.io/gorm"

	"github.com/Mauricio-Carrion/LeadMeBackend/models/companie"
)

func NewCompanieController(newCompanie types.NewCompanie) (*gorm.DB, error) {
	if newCompanie.Name == "" || newCompanie.Razao == "" || newCompanie.Cpf_cnpj == "" {
		return nil, gorm.ErrInvalidData
	}

	data, err := companie.NewCompanieModel(&newCompanie)

	if err != nil {
		return nil, err
	}

	return data, nil
}