package main

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func (server *server) saveBookingDetail(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "saveBookingDetail",
			"subFunc": "ioutil.ReadAll",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the request
	msg, err := server.nats.Request("data", data, time.Second)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "saveBookingDetail",
			"subFunc": "server.nats.Request",
		}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(msg.Data)
	return
}
