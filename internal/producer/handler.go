package producer

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aakash19here/email-microservices/internal/broker"
	"github.com/aakash19here/email-microservices/internal/email"
	"github.com/google/uuid"
)

type Handler struct{ Broker *broker.Broker }

func (h *Handler) HandleEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var e email.Email

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := e.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e.ID = uuid.NewString()
	e.CreatedAt = time.Now().UTC()

	body, _ := e.Marshal()

	if err := h.Broker.Publish(r.Context(), body); err != nil {
		http.Error(w, "failed to queue email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted) // 202: accepted, not yet sent
	json.NewEncoder(w).Encode(map[string]string{"id": e.ID, "status": "queued"})
}
