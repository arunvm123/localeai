package main

import (
	"flag"
	"net/http"

	"github.com/arunvm/locale/config"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

type server struct {
	nats *nats.Conn
}

func newServer() *server {
	s := server{}
	return &s
}

func main() {
	server := newServer()

	log.SetFormatter(&log.JSONFormatter{})

	// Reading file path from flag
	filePath := flag.String("config-path", "config.yaml", "filepath to configuration file")
	flag.Parse()

	// Reading config variables
	config, err := config.Initialise(*filePath)
	if err != nil {
		log.Fatalf("Failed to read config\n%v", err)
	}

	// Message broker
	server.nats, err = nats.Connect(config.NatsURL)
	if err != nil {
		log.Fatalf("Failed to connect to  nats server\n%v", err)
	}
	defer server.nats.Close()

	http.HandleFunc("/save/booking/detail", server.saveBookingDetail)

	log.Println("Server listening on port ", config.Port)
	http.ListenAndServe(":"+config.Port, nil)
}
