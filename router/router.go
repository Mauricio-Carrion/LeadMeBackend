package router

import (
	"fmt"

	"github.com/Mauricio-Carrion/LeadMeBackend/configs"

	"github.com/Mauricio-Carrion/LeadMeBackend/routes"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	routes.Welcome(router)
	routes.ConnectWhatsmeow(router)
	routes.CreateUser(router)
	router.Run(fmt.Sprintf(":%s", configs.GetServerPort()))
}