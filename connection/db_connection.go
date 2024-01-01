package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mauricio-Carrion/LeadMeBackend/configs"
	_ "github.com/lib/pq"
)

func DBConnection() (*sql.DB, error){
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	conf.Host, 
	conf.Port, 
	conf.User, 
	conf.Password, 
	conf.Database)
	
	conn, err := sql.Open("postgres", sc)
	if err != nil{
		log.Fatal(err)
	}

	err = conn.Ping()

	return conn, err
}