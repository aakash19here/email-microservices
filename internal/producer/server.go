package producer

import (
	"context"
	"net/http"
	"time"
)

func NewServer(port string, h *Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /emails", h.HandleEmail)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	})

	return &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func Shutdown(s *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_ = s.Shutdown(ctx)
}
