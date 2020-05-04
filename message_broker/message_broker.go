package msgbroker

import (
	"log"

	"github.com/nats-io/nats.go"
)

func Subscribe(nats *nats.Conn, subject string, callback func(msg *nats.Msg)) (*nats.Subscription, error) {
	sub, err := nats.Subscribe(subject, callback)
	if err != nil {
		log.Fatalf("Error in message broker\n%v", err)
		return nil, err
	}

	return sub, nil
}
