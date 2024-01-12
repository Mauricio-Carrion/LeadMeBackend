package routes

import (
	"net/http"

	"github.com/Mauricio-Carrion/LeadMeBackend/types"

	controllers "github.com/Mauricio-Carrion/LeadMeBackend/controllers/user"
	"github.com/gin-gonic/gin"
)


func CreateUser(router *gin.Engine) {
	router.POST("/user", func(ctx *gin.Context) {
		var user types.NewUser
		  
		if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
		

		err := controllers.NewUserController(user)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(200, gin.H{
			"message": "User created",
		})
	})
}