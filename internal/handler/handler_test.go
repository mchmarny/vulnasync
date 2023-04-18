package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mchmarny/vulnasync/internal/pubsub"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	ca "google.golang.org/api/containeranalysis/v1"
	api "google.golang.org/api/pubsub/v1"
)

func TestHandler(t *testing.T) {
	provider = testProvider
	sender = testSender

	d := []byte(`{
		"name": "projects/test/occurrences/d2342144-8a7e-4f3c-b3ba-87ebbe3ac72d",
		"kind": "VULNERABILITY", 
		"notificationTime": "2023-03-30T23:09:28.471565Z"
	}`)

	m := pubsub.Message{
		Message: api.PubsubMessage{
			Data: base64.StdEncoding.EncodeToString(d),
		},
	}

	j, err := json.Marshal(m)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(j))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Process)

	handler.ServeHTTP(rr, req)
	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
}

func testSender(_ context.Context, occ *ca.Occurrence) error {
	if occ == nil {
		return errors.New("occurrence is nil")
	}

	if err := json.NewEncoder(os.Stdout).Encode(occ); err != nil {
		return errors.Wrap(err, "failed to encode occurrence")
	}

	return nil
}

func testProvider(_ context.Context, name string) (*ca.Occurrence, error) {
	b, err := os.ReadFile("test.json")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read test file: %s", name)
	}

	var occ ca.Occurrence
	if err := json.Unmarshal(b, &occ); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal test file: %s", name)
	}

	return &occ, nil
}
