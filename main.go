package main

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/configs"
	"github.com/Mauricio-Carrion/LeadMeBackend/db"
	"github.com/Mauricio-Carrion/LeadMeBackend/router"
	"go.mau.fi/whatsmeow/store"
	"google.golang.org/protobuf/proto"
)

func main() {
	store.DeviceProps.Os = proto.String("LeadMe")
	configs.InitConfig()
	db.DBConnection()
	router.Router()
}