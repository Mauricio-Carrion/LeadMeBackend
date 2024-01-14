package user

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Mauricio-Carrion/LeadMeBackend/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

	"github.com/Mauricio-Carrion/LeadMeBackend/models/user"
)

func LoginUserController(login *types.LoginData) (*types.TokenType, error) {

	if login.Email == "" || login.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := user.LoginUserModel(login)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = godotenv.Load(".env")

	if err != nil {
			log.Fatal("Error loading .env file")
	}

	requiredToken := os.Getenv("API_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.Uuid
	claims["companie_uuid"] = user.Companie_uuid
	claims["name"] = user.Name
	claims["admin"] = user.IsAdmin
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()


	t, err := token.SignedString([]byte(requiredToken))
	
	if err != nil {
			return nil, err
	}
	
	return &types.TokenType{
		AccessToken: t,
	}, nil
}