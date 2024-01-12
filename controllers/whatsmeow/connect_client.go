package controllers

import (
	"context"
	"log"
	"os"

	"github.com/Mauricio-Carrion/LeadMeBackend/connection"
	models "github.com/Mauricio-Carrion/LeadMeBackend/models/whatsmeow"
	"github.com/gin-gonic/gin"
	"github.com/mdp/qrterminal/v3"
)

func ConnectClient(c *gin.Context) {
	conn := connection.CreateWhatsmeowContainer()
	client := models.WhatsmeowNewClient(conn)
	
	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		
		if err != nil {
			log.Fatal(err)
		}

		for evt := range qrChan {
			if evt.Event == "code" { 
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				c.JSON(200, gin.H{
					"QR Code": evt.Code,
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