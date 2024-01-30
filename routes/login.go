package routes

import (
	"net/http"

	"github.com/Mauricio-Carrion/LeadMeBackend/controllers/user"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"github.com/gin-gonic/gin"
)

func Login(router *gin.Engine) {
	router.POST("/login", func(ctx *gin.Context) {
		var newLogin types.LoginData
		  
		if err := ctx.ShouldBindJSON(&newLogin); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
		
		data, err := user.LoginUserController(&newLogin)

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

		ctx.JSON(200, 
			 data,
		)
	})
}