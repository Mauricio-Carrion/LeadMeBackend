package user

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"gorm.io/gorm"

	"github.com/Mauricio-Carrion/LeadMeBackend/models/user"
)

func NewUserController(newUser types.NewUser) (*gorm.DB, error) {
	if newUser.Name == "" || newUser.Email == "" || newUser.Phone == "" || newUser.Password == "" {
		return nil, gorm.ErrInvalidData
	}
	
	data ,err := user.NewUserModel(&newUser)

	if err != nil {
		return nil, err
	}

	return data, nil
}