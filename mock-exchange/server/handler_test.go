package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchDataHandlerNeedsToBePost(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fetch", nil)
	if err != nil {
		t.Fatalf("Error calling handler: %s", err)
	}

	rr := httptest.NewRecorder()
	handler := FetchDataHandler{}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestWeFetchDataWeWantToDisplay(t *testing.T) {
	requestData := strings.NewReader(`{"from":"USD","to":"BTC"}`)

	req, err := http.NewRequest(http.MethodPost, "/fetch", requestData)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	rr := httptest.NewRecorder()
	handler := FetchDataHandler{}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}
