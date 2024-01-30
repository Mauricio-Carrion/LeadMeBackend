package user

import (
	"errors"

	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
)

func GetJidModel(userUuid string) (*types.JidResponse, error) {
	var result types.JidResponse
	queryResult := db.DBConnection().Model(&db.User{}).Select("jid").Where("uuid = ?", userUuid).Find(&result)

	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	if queryResult.RowsAffected == 0 {
		return nil, errors.New(" not found")
	}

	return &result, nil
}