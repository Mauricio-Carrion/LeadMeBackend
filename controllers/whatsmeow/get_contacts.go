package controllers

import (
	"errors"

	models "github.com/Mauricio-Carrion/LeadMeBackend/models/whatsmeow"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
)

func GetContacts(jid string, limit string, offset string) ([]*types.WhatsmeowContacts, error) {
	if jid == "" {
		return nil , errors.New("jid is required")
	}
	
	data, err := models.WhatsmeowContacts(jid, limit, offset)

	if err != nil {
		return nil, err
	}

	return data, nil
}