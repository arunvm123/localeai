package main

import (
	"encoding/json"

	"github.com/arunvm/locale/models"
	"github.com/go-playground/validator"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

type response struct {
	ID    int    `json:"id"`
	Type  int    `json:"type"`
	Error string `json:"error"`
}

const (
	success = 1
	fail    = 2
)

func (server *server) consumer() func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var bd models.BookingDetail

		err := json.Unmarshal(msg.Data, &bd)
		if err != nil {
			log.WithFields(log.Fields{
				"func":    "consumer",
				"subFunc": "json.Unmarshal",
			}).Error(err)
			handleResponse(msg, 0, fail, err.Error())
			return
		}

		err = validator.New().Struct(&bd)
		if err != nil {
			log.WithFields(log.Fields{
				"func":    "consumer",
				"subFunc": "validator.New().Struct",
				"id":      bd.ID,
			}).Error(err)
			handleResponse(msg, bd.ID, fail, err.Error())
			return

		}

		err = models.SaveBookingDetail(server.db, &bd)
		if err != nil {
			log.WithFields(log.Fields{
				"func":    "consumer",
				"subFunc": "models.SaveBookingDetail",
				"id":      bd.ID,
			}).Error(err)
			handleResponse(msg, bd.ID, fail, err.Error())
			return
		}

		handleResponse(msg, bd.ID, success, "")
		return
	}
}

func handleResponse(msg *nats.Msg, id int, responseType int, errorMessage string) {
	resp := response{
		ID:    id,
		Error: errorMessage,
		Type:  responseType,
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "handleResponse",
			"subFunc": "json.Marshal",
			"id":      resp.ID,
		}).Error(err)
		return
	}

	err = msg.Respond(data)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "handleResponse",
			"subFunc": "msg.Respond",
			"id":      resp.ID,
		}).Error(err)
		return
	}

	return
}
