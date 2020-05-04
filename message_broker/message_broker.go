package msgbroker

import (
	log "github.com/sirupsen/logrus"

	"github.com/nats-io/nats.go"
)

func Subscribe(nats *nats.Conn, subject string, callback func(msg *nats.Msg)) (*nats.Subscription, error) {
	sub, err := nats.Subscribe(subject, callback)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "Subscribe",
			"subFunc": "nats.Subscribe",
			"subject": subject,
		}).Error(err)
		return nil, err
	}

	return sub, nil
}
