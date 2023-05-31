package main

import (
	"fmt"
	"log"

	config "github.com/auth/service/pkg/config"
	db "github.com/auth/service/pkg/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load the Congiguration File: ", err)
		return
	}
	DB, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatalln("Failed to connect to the Database: ", err)
		return
	}
	fmt.Println(DB)

}
