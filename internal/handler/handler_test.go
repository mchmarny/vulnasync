package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mchmarny/vulnasync/internal/pubsub"
	"github.com/stretchr/testify/assert"
	api "google.golang.org/api/pubsub/v1"
)

func TestHandler(t *testing.T) {
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
