package routes

import (
	"net/http"

	"github.com/Mauricio-Carrion/LeadMeBackend/controllers/companie"
	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"github.com/gin-gonic/gin"
)

func CreateCompanie(router *gin.Engine) {
router.POST("/companie", func(ctx *gin.Context) {
		var newCompanie types.NewCompanie
		  
		if err := ctx.ShouldBindJSON(&newCompanie); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
		
		data, err := companie.NewCompanieController(newCompanie)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})

			return
		}

		ctx.JSON(200, gin.H{
			"message": "Companie created",
			"data":    data,
		})
	})
}