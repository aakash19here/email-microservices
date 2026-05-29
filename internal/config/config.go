package config

import (
	"os"
	"time"
)

type Config struct {
	RabbitURL string
	Queue     string
	HTTPPort  string
	SendDelay time.Duration
	Prefetch  int
}

func Load() Config {
	return Config{
		RabbitURL: env("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		Queue:     env("QUEUE_NAME", "emails"),
		HTTPPort:  env("HTTP_PORT", "8080"),
		SendDelay: 3 * time.Second,
	}
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func envDuration(k string, def time.Duration) time.Duration {
	if v := os.Getenv(k); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}
