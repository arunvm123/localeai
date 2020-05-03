package main

import (
	"flag"
	"sync"

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

	msgbroker.Subscribe(server.nats, "data", server.consumer())

	log.Println("Server subscribed to queue")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
