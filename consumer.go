package main

import (
	"encoding/json"

	"github.com/arunvm/locale/models"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func (server *server) consumer() func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var bd models.BookingDetail

		err := json.Unmarshal(msg.Data, &bd)
		if err != nil {
			log.WithFields(log.Fields{
				"subFunc": "json.Unmarshal",
			}).Error(err)
			return
		}

		err = models.CreateBookingDetail(server.db, &bd)
		if err != nil {
			log.WithFields(log.Fields{
				"subFunc": "models.CreateBookingDetail",
			}).Error(err)
			return
		}

		err = msg.Respond([]byte("Message recieved sucessfully"))
		if err != nil {
			log.WithFields(log.Fields{
				"subFunc": "msg.Respond",
			}).Error(err)
			return
		}

		log.Printf("Message: %v", bd)
	}
}
