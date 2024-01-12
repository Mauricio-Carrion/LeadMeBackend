package db

import (
	"fmt"
	"log"

	"github.com/Mauricio-Carrion/LeadMeBackend/configs"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Active   bool
	IsAdmin  bool
	Jid      string
}

func DBConnection() *gorm.DB {
	pgconn := postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	configs.GetDB().Host, 
	configs.GetDB().Port, 
	configs.GetDB().User, 
	configs.GetDB().Password, 
	configs.GetDB().Database))

	db, err := gorm.Open(pgconn, &gorm.Config{})
  
	if err != nil {
    log.Fatal("failed to connect database")

  }

	db.AutoMigrate(
		&User{}, 
	)

	return db
}

 