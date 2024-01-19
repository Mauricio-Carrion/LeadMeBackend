package models

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func WhatsmeowNewClient(container *sqlstore.Container) *whatsmeow.Client {
	deviceStore := container.NewDevice()

	

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	
	return client
}