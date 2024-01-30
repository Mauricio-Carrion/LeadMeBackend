package user

import (
	"errors"

	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
)

func LoginUserModel(login *types.LoginData) (*types.SelectUser, error) {
	var result types.UserResponse
	
	queryResult := db.DBConnection().Model(&db.User{}).Where("email = ?", login.Email).Find(&result)

	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	if queryResult.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &types.SelectUser{
		Uuid: result.Uuid, 
		Name: result.Name, 
		Email: result.Email, 
		Active: result.Active, 
		IsAdmin: result.IsAdmin, 
		Companie_uuid: result.Companie_uuid,
		},
		nil
}