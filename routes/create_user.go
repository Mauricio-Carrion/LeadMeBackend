package routes

import (
	"net/http"

	"github.com/Mauricio-Carrion/LeadMeBackend/types"

	"github.com/Mauricio-Carrion/LeadMeBackend/controllers/user"
	"github.com/gin-gonic/gin"
)


func CreateUser(router *gin.Engine) {
	router.POST("/user", func(ctx *gin.Context) {
		var newUser types.NewUser
		  
		if err := ctx.ShouldBindJSON(&newUser); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
		
		data, err := user.NewUserController(newUser)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}

		ctx.JSON(200, gin.H{
			"message": "User created",
			"data":    data,
		})
	})
}