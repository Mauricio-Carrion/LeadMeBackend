package routes

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/controllers"
	"github.com/gin-gonic/gin"
)

func ConnectWhatsmeow(router *gin.Engine) {
		router.GET("/connect", controllers.ConnectClient) 
}