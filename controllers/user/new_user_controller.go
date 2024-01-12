package user

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
)

func NewUserController(newUser types.NewUser) error {
	if newUser.Name == "" || newUser.Email == "" || newUser.Phone == "" || newUser.Password == "" {
		return nil
	}
	
	db.DBConnection()

	return nil
}