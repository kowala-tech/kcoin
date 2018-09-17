package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchDataHandlerNeedsToBePost(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fetch", nil)
	if err != nil {
		t.Fatalf("Error calling handler: %s", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fetchDataHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}
