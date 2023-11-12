package application

import (
	"encoding/json"
	"log"

	"example.com/hexagonal-listener/src/domain"
)

type MessageEventHandler struct{}

func (h *MessageEventHandler) HandleEvent(event *domain.Event) error {

	log.Printf("Event received: ID=%s, Payload=%s \n", event.ID, event.Payload)

	var transaction domain.Transaction

	err := json.Unmarshal([]byte(event.Payload), &transaction)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	//log.Printf("transactionNumber: %s", transaction.Data.TransactionNumber)

	token, err := GetToken()

	if err != nil {
		log.Printf("Error Token: %v", err)
	}

	data, err := GetTransaction(token, transaction)

	if err != nil {
		log.Printf("Error GET %v", err)
	}

	log.Printf("GET transaction %s -> %s", event.ID, data)

	return nil
}
