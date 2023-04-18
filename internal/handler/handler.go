package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mchmarny/vulnasync/internal/aa"
	"github.com/mchmarny/vulnasync/internal/pubsub"
	"github.com/mchmarny/vulnasync/internal/sender/stdout"
	"github.com/rs/zerolog/log"
	ca "google.golang.org/api/containeranalysis/v1"
)

const (
	messageTypeVulnerability = "VULNERABILITY"
)

var (
	// default occurrence sender
	sender OccurrenceSender = stdout.Sender
)

// OccurrenceSender is a function that sends an occurrence to a specific.
type OccurrenceSender func(ctx context.Context, occ *ca.Occurrence) error

// Process handles the incoming Pub/Sub message.
func Process(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("processing...")

	if r.Method != http.MethodPost {
		log.Printf("invalid method: %s", r.Method)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	v, err := pubsub.GetMessageDate[aa.VulnerabilityMessage](r)
	if err != nil {
		log.Printf("error while getting message data: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if v.Kind != messageTypeVulnerability {
		log.Printf("ignoring message of kind: %s", v.Kind)
		fmt.Fprintf(w, "invalid kind: %s", v.Kind)
		return
	}

	occ, err := aa.GetOccurrence(r.Context(), v.Name)
	if err != nil {
		log.Printf("error while getting occurrence: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := sender(r.Context(), occ); err != nil {
		log.Printf("error while sending occurrence: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "ok")
}
