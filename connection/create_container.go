package connection

import (
	"fmt"
	"log"

	"github.com/Mauricio-Carrion/LeadMeBackend/configs"
	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func CreateWhatsmeowContainer() (*sqlstore.Container ){
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	
	container, err := sqlstore.New("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.GetDB().Host, 
		configs.GetDB().Port, 
		configs.GetDB().User, 
		configs.GetDB().Password, 
		configs.GetDB().Database), 
		dbLog)
	
	if err != nil {
		log.Fatal(err)
	}

	return container
}
