package routes

import (
	controllers "github.com/Mauricio-Carrion/LeadMeBackend/controllers/whatsmeow"
	"github.com/gin-gonic/gin"
)

func ConnectWhatsmeow(router *gin.Engine) {
		router.GET("/connect", controllers.ConnectClient) 
}