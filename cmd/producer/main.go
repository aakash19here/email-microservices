package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aakash19here/email-microservices/internal/broker"
	"github.com/aakash19here/email-microservices/internal/config"
	"github.com/aakash19here/email-microservices/internal/producer"
)

func main() {
	cfg := config.Load()

	b, err := broker.Connect(cfg.RabbitURL, cfg.Queue)

	if err != nil {
		log.Fatal(err)
	}

	defer b.Close()

	srv := producer.NewServer(cfg.HTTPPort, &producer.Handler{Broker: b})

	go func() {
		log.Printf("producer listening on :%s", cfg.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	producer.Shutdown(srv)

}
