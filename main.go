package main

import (
	"log"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = NewWSServer(config.Server)

	if err != nil {
		log.Fatal(err.Error())
	}
}
