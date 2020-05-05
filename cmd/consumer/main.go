package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	msgbroker "github.com/arunvm/locale/message_broker"

	log "github.com/sirupsen/logrus"

	"github.com/arunvm/locale/config"
	"github.com/arunvm/locale/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nats-io/nats.go"
)

type server struct {
	db   *gorm.DB
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

	// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	server.db, err = gorm.Open("postgres", "host= "+config.Database.Host+" port="+config.Database.Port+" user="+config.Database.User+" dbname="+config.Database.DatabaseName+" password="+config.Database.Password+" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database\n%v", err)
	}

	server.db.LogMode(true)
	models.MigrateDB(server.db)

	// Message broker
	server.nats, err = nats.Connect(config.NatsURL)
	if err != nil {
		log.Fatalf("Failed to connect to  nats server\n%v", err)
	}
	defer server.nats.Close()

	sub, err := msgbroker.Subscribe(server.nats, "data", server.consumer())
	if err != nil {
		log.Fatalf("Failed to subscribe to message broker\n%v", err)
	}

	log.Println("Subscribed to message broker")

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan

	log.Println("Terminating...")

	err = sub.Drain()
	if err != nil {
		log.Fatalf("Error when draining message broker\n%v", err)
	}
}
