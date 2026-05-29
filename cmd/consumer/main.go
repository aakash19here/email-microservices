package main

import (
	"log"

	"github.com/aakash19here/email-microservices/internal/broker"
	"github.com/aakash19here/email-microservices/internal/config"
	"github.com/aakash19here/email-microservices/internal/consumer"
)

func main() {
	cfg := config.Load()

	b, err := broker.Connect(cfg.RabbitURL, cfg.Queue)

	if err != nil {
		log.Fatal(err)
	}

	defer b.Close()

	if err := consumer.Run(b, cfg.SendDelay); err != nil {
		log.Fatal(err)
	}
}
