package utils

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func WhatsmeowEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}