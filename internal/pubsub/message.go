package pubsub

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	api "google.golang.org/api/pubsub/v1"
)

// Message is the payload of a Pub/Sub event.
type Message struct {
	Message      api.PubsubMessage `json:"message"`
	Subscription string            `json:"subscription"`
}

// GetMessageDate decodes the message data and returns the specified type.
func GetMessageDate[T any](r *http.Request) (*T, error) {
	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return nil, errors.Wrap(err, "failed to decode message")
	}

	d, err := base64.StdEncoding.DecodeString(m.Message.Data)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding message data")
	}

	var e T
	if err = json.Unmarshal(d, &e); err != nil {
		return nil, errors.Wrap(err, "error parsing event")
	}
	return &e, nil
}
