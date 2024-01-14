package user

import (
	"errors"
	"fmt"

	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
)

func LoginUserModel(login *types.LoginData) (*types.SelectUser, error) {
	var result types.User
	
	queryResult := db.DBConnection().Where("email = ?", login.Email).Find(&result)

	fmt.Println(result)

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