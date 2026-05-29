package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitURL string
	Queue     string
	HTTPPort  string
	SendDelay time.Duration
	Prefetch  int
	ResendKey string
	From      string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		RabbitURL: env("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		Queue:     env("QUEUE_NAME", "emails"),
		HTTPPort:  env("HTTP_PORT", "8080"),
		ResendKey: env("RESEND_API_KEY", "re_xxxxxxxxx"),
		From:      env("FROM_EMAIL", "onboarding@resend.dev"),
		SendDelay: 3 * time.Second,
	}
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
