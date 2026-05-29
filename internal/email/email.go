package email

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
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
