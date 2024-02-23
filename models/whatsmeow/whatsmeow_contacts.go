package models

import (
	"errors"

	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
)


func WhatsmeowContacts(jid string, limit string, offset string ) ([]*types.WhatsmeowContacts, error){
		var result []*types.WhatsmeowContacts
		
		queryResult := 
		  db.DBConnection().Raw("SELECT DISTINCT their_jid, first_name, full_name, push_name, business_name FROM whatsmeow_contacts WHERE our_jid like	'%" + jid + "%' AND their_jid NOT LIKE '%" + jid + "%' ORDER BY full_name LIMIT " + limit + " OFFSET " + offset + "").Scan(&result)
		
		if queryResult.Error != nil {
			return nil, queryResult.Error
		}

		if queryResult.RowsAffected == 0 {
			return nil, errors.New("not found")
		}

		return result, nil
}