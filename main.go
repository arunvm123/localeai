package main

import (
	"flag"
	"log"

	"github.com/arunvm/locale/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type server struct {
	db *gorm.DB
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
		panic(err)
	}
}
