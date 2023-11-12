package application

import (
	"context"

	"example.com/hexagonal-listener/src/domain"
)

type PubSubService interface {
	SubscribeToEvents(ctx context.Context, handler EventHandler) error
}

type EventHandler interface {
	HandleEvent(event *domain.Event) error
}
