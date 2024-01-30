package routes

import (
	"fmt"

	"github.com/Mauricio-Carrion/LeadMeBackend/controllers/user"
	"github.com/gin-gonic/gin"
)


func GetJid(router *gin.Engine) {
	router.GET("/user/jid/:uuid", func(ctx *gin.Context) {

		  
		// if err := ctx.ShouldBindJSON(&); err != nil {
    //     ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    //     return
    // }

		userUuid := ctx.Param("uuid")

		fmt.Println(userUuid)

		if userUuid == "" {
			ctx.JSON(400, gin.H{
				"message": "user uuid is required",
			})
		  return
		}
		
		data, err := user.GetJidController(userUuid)

		if err != nil {
			if err.Error() == "invalid credentials" || err.Error() == "email and password are required" {
				ctx.JSON(401, gin.H{
					"message": err.Error(),
				})
			}

			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}

		ctx.JSON(200, gin.H{
			"jid":    data,
		})
	})
}