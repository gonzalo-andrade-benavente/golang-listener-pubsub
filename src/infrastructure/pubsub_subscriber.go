package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"example.com/hexagonal-listener/src/application"
	"example.com/hexagonal-listener/src/domain"
)

type PubSubSubscriber struct {
	Subscription *pubsub.Subscription
}

func NewPubSubSubscriber(subscription *pubsub.Subscription) *PubSubSubscriber {
	return &PubSubSubscriber{Subscription: subscription}
}

func (s *PubSubSubscriber) SubscribeToEvents(ctx context.Context, handler application.EventHandler) error {

	err := s.Subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {

		event := &domain.Event{
			ID:      msg.Attributes["eventId"],
			Payload: string(msg.Data),
		}

		if err := handler.HandleEvent(event); err != nil {
			log.Printf("Error to handle message %v", err)
		}

		//msg.Ack()

	})

	return err

}
