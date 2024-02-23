package db

import (
	"fmt"
	"log"

	"github.com/Mauricio-Carrion/LeadMeBackend/configs"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Uuid     string
	Name     string
	Email    string
	Password string
	Active   bool
	IsAdmin  bool
	Jid      string
	Companie_uuid uuid.UUID
	Companie Companie `gorm:"foreignKey:Companie_uuid"`
}

type Companie struct {
	Uuid     string `gorm:"primaryKey"`
	Name     string
	Razao    string
	Cpf_cnpj string
	Active   bool
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
		&Companie{}, 
	)

	return db
}

 