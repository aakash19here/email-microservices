package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	conn  *amqp.Connection
	queue string
	ch    *amqp.Channel
}

func Connect(url, queue string) (*Broker, error) {
	var conn *amqp.Connection
	var err error

	for i := range 10 {
		conn, err = amqp.Dial(url)

		if err == nil {
			break
		}

		log.Printf("rabbitmq not ready (%d/10): %v url is %s", i+1, err, url)

		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("dial rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("open channel: %w", err)
	}

	if _, err := ch.QueueDeclare(queue, true, false, false, false, nil); err != nil {
		return nil, fmt.Errorf("declare queue %w", err)
	}

	return &Broker{conn: conn, queue: queue, ch: ch}, nil
}

func (b *Broker) Publish(ctx context.Context, body []byte) error {
	return b.ch.PublishWithContext(ctx, "", b.queue, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	})
}

func (b *Broker) Consume() (<-chan amqp.Delivery, error) {
	return b.ch.Consume(b.queue, "", false, false, false, false, nil)
}

func (b *Broker) Close() {
	if b.ch != nil {
		_ = b.ch.Close()
	}
	if b.conn != nil {
		_ = b.conn.Close()
	}
}
