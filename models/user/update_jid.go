package user

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"gorm.io/gorm"
)

func UpdateJID(userUuid string,jid string) ( *gorm.DB, error) {
		queryResult := db.DBConnection().Model(&db.User{}).Where("uuid = ?", userUuid).Update("jid", jid)

		if queryResult.Error != nil {
			return nil, queryResult.Error
		}

		return queryResult, nil
}