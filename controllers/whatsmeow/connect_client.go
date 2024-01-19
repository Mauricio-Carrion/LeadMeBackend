package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Mauricio-Carrion/LeadMeBackend/connection"
	"github.com/Mauricio-Carrion/LeadMeBackend/models/user"
	models "github.com/Mauricio-Carrion/LeadMeBackend/models/whatsmeow"
	"github.com/gin-gonic/gin"
	"github.com/mdp/qrterminal/v3"
)


func ConnectWPClient(c *gin.Context) {
	conn := connection.CreateWhatsmeowContainer()
	client := models.WhatsmeowNewClient(conn)

	go func ()  {
		for{
			if client.Store.ID != nil {
				 _ ,err :=	user.UpdateJID(
						c.Keys["uuid"].(string),
						client.Store.ID.String(),
					)

				fmt.Println(client.Store.Platform)
				fmt.Println(client.Store)

				 if err != nil {
				 	log.Fatal(err)
				 }
				break
			}
	  } 	
	}()

	if client.Store.ID == nil {
			qrChan, _ := client.GetQRChannel(context.Background())
			err := client.Connect()

			if err != nil {
				log.Fatal(err)
			}
		
			for evt := range qrChan {
			
				if evt.Event == "code" {
					qrterminal.Generate(evt.Code, qrterminal.L, os.Stdout)
				  c.JSON(200, gin.H{
				  "qrCode": evt.Code,
			    })
				 break
				}
			}

			} else {
				err := client.Connect()

				if err != nil {
					log.Fatal(err)
				}
			}
}