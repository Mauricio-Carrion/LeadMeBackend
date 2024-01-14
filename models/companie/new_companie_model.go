package companie

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewCompanieModel(newCompanie *types.NewCompanie) (*gorm.DB, error) {

	result := db.DBConnection().Create(&db.Companie{
		Uuid:     uuid.New(),
		Name:     newCompanie.Name,
		Razao:    newCompanie.Razao,
		Cpf_cnpj: newCompanie.Cpf_cnpj,
		Active:   newCompanie.Active,
	})

	if result.Error != nil {
		return result, result.Error
	}

	return result, nil
}