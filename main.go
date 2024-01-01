package main

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/configs"
	"github.com/Mauricio-Carrion/LeadMeBackend/router"
)

func main() {
	configs.InitConfig()
	router.Router()
}