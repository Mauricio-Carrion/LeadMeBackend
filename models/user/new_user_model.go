package user

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserModel(newUser *types.NewUser) ( *gorm.DB, error) {

	result := db.DBConnection().Create(&db.User{
		Uuid:     uuid.New(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Active:   newUser.Active,
		IsAdmin:  newUser.IsAdmin,
		Companie_uuid: newUser.Companie_uuid,
	})

	if result.Error != nil {
		return result, result.Error
	}

	return result, nil
}