package email

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aakash19here/email-microservices/internal/config"
	"github.com/resend/resend-go/v3"
)

type Email struct {
	ID        string    `json:"id"`
	To        string    `json:"to"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (e Email) Validate() error {
	if strings.TrimSpace(e.To) == "" {
		return errors.New("To is required")
	}

	if strings.TrimSpace(e.Subject) == "" {
		return errors.New("Subject is required")
	}

	return nil
}

func (e Email) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func Unmarshal(b []byte) (Email, error) {
	var e Email

	err := json.Unmarshal(b, &e)

	return e, err
}

func (e *Email) Send(body, to string) error {
	cfg := config.Load()
	client := resend.NewClient(cfg.ResendKey)

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("Acme <%s>", cfg.From),
		To:      []string{to},
		Html:    fmt.Sprintf("<strong>%s</strong>", body),
		Subject: "Hello from Golang",
	}

	_, err := client.Emails.Send(params)

	return err
}
