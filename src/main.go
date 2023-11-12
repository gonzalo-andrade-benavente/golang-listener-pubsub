package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"example.com/hexagonal-listener/src/application"
	"example.com/hexagonal-listener/src/infrastructure"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	subscription := client.Subscription(os.Getenv("GCP_SUBSCRIPTION_ID"))

	subscriber := infrastructure.NewPubSubSubscriber(subscription)

	eventHandler := &application.MessageEventHandler{}

	err = subscriber.SubscribeToEvents(ctx, eventHandler)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
