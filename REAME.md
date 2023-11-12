export GOOGLE_APPLICATION_CREDENTIALS=/home/gonzalo/golang/hexagonal-listener/key.json

src/
|-- application
|   |-- pubsub_service.go
|   |-- event_handler.go
|-- domain
|   |-- event.go
|-- infrastructure
|   |-- pubsub_publisher.go
|   |-- pubsub_subscriber.go
|-- main.go