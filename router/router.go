package router

import (
	"fmt"

	"github.com/Mauricio-Carrion/LeadMeBackend/auth"
	"github.com/Mauricio-Carrion/LeadMeBackend/configs"

	"time"

	"github.com/Mauricio-Carrion/LeadMeBackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))
	router.Use(gin.Logger())
	routes.Welcome(router)
	routes.Login(router)
	router.Use(auth.AuthMiddleware())
	routes.ConnectWhatsmeow(router)
	routes.CreateUser(router)
	routes.CreateCompanie(router)
	routes.GetJid(router)
	router.Run(fmt.Sprintf(":%s", configs.GetServerPort()))
}