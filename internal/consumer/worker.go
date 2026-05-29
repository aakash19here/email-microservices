package consumer

import (
	"log"
	"time"

	"github.com/aakash19here/email-microservices/internal/broker"
	"github.com/aakash19here/email-microservices/internal/email"
)

func Run(b *broker.Broker, delay time.Duration) error {
	msgs, err := b.Consume()

	if err != nil {
		return err
	}

	log.Println("consumer waiting for messages...")

	for d := range msgs {
		e, err := email.Unmarshal(d.Body)

		if err != nil {
			log.Printf("bad message, dropping: %v", err)

			_ = d.Nack(false, false) // don't requeue malformed messages

			continue
		}

		log.Printf("sending email %s to %s ...", e.ID, e.To)
		time.Sleep(delay) // artificial delay simulating SMTP send
		log.Printf("sent email %s", e.ID)

		_ = d.Ack(false) // ack only after the send "succeeds"
	}

	return nil
}
