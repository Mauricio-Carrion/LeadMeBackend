package user

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"github.com/Mauricio-Carrion/LeadMeBackend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserModel(newUser *types.NewUser) ( *gorm.DB, error) {
	hashedPassword, err := utils.EncryptPassword(newUser.Password)

	if err != nil {
		return nil, err
	}

	result := db.DBConnection().Create(&db.User{
		Uuid:     uuid.New().String(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: hashedPassword,
		Active:   newUser.Active,
		IsAdmin:  newUser.IsAdmin,
		Companie_uuid: newUser.Companie_uuid,
	})

	if result.Error != nil {
		return result, result.Error
	}

	return result, nil
}