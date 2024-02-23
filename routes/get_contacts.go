package routes

import (
	controllers "github.com/Mauricio-Carrion/LeadMeBackend/controllers/whatsmeow"
	"github.com/gin-gonic/gin"
)

func GetContacts(router *gin.Engine) {
	router.GET("/contacts/:jid", func(ctx *gin.Context) {
		
		data, err := controllers.GetContacts(ctx.Param("jid"), ctx.Query("contacts"), ctx.Query("page"))

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}

		ctx.JSON(200, gin.H{
			"contacts":    data,
		})
	})
}