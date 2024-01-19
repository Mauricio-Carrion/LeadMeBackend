package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func AuthMiddleware() gin.HandlerFunc {

    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    requiredToken := os.Getenv("API_SECRET")

    if requiredToken == "" {
        log.Fatal("API_SECRET must be set")
    }

		fmt.Println(requiredToken)

    return func(c *gin.Context) {
	
				jwtToken := c.Request.Header.Get("Authorization")
				claims := jwt.MapClaims{}
	
				if jwtToken == "" {
					c.JSON(http.StatusBadRequest, gin.H{"message": "Token required"})
					c.Abort()
					return
				}else {
					jwtToken = strings.Replace(jwtToken, "Bearer ", "", -1)
				}

				token, err := jwt.ParseWithClaims(jwtToken,claims,func(token *jwt.Token) (interface{}, error) {
         return []byte(requiredToken), nil
        })

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
					c.Abort()
					return
				}

				if !token.Valid {
					c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid token"})
					c.Abort()
					return
				}
				
				c.Keys = map[string]interface{ }{
					"uuid": claims["uuid"],
					"email": claims["email"],
					"name": claims["name"],
					"admin": claims["admin"],
					"companie_uuid": claims["companie_uuid"],
				}
        c.Next()
    }
}