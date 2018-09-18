package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/kowala-tech/kcoin/mock-exchange/app"

	"github.com/patrickmn/go-cache"
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

func TestWeFetchDataWeWantToDisplayAndIsSavedToCache(t *testing.T) {
	requestData := strings.NewReader(
		`{
			"sell":[
				{"amount":0.358, "rate":6326.83689418},
				{"amount":0.1427, "rate":6326.83689421}
			], 
			"buy":[
				{"amount":0.0021, "rate":6214.3034165},
				{"amount":0.0029, "rate":6203.01833171}
			]
		}`,
	)

	req, err := http.NewRequest(http.MethodPost, "/fetch", requestData)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	c := cache.New(5*time.Minute, 10*time.Minute)

	rr := httptest.NewRecorder()
	handler := FetchDataHandler{c}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	savedRequest, ok := c.Get(app.CacheRequestKey)
	if !ok {
		t.Fatalf("Failed to assert that request was saved in the cache.")
	}

	expectedRequest := app.Request{
		Sell: []app.RateValue{
			{
				Amount: 0.358,
				Rate:   6326.83689418,
			},
			{
				Amount: 0.1427,
				Rate:   6326.83689421,
			},
		},
		Buy: []app.RateValue{
			{
				Amount: 0.0021,
				Rate:   6214.3034165,
			},
			{
				Amount: 0.0029,
				Rate:   6203.01833171,
			},
		},
	}

	assert.Equal(t, expectedRequest, savedRequest)
}

func TestGetRatesHandler_ItFailsIfItsNotGetRequest(t *testing.T) {
	handler := GetRatesHandler{}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/req", nil)
	if err != nil {
		t.Fatalf("Failed to create Get Rates request: %s", err)
	}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestGetRatesHandler_ItFailsIfItWeDidNotCallFetchBefore(t *testing.T) {
	c := cache.New(5*time.Minute, 10*time.Minute)

	handler := GetRatesHandler{
		cache:       c,
		transformer: nil,
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/exrates", nil)
	if err != nil {
		t.Fatalf("Failed to create request for get rates handler.")
	}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	errorResponse := ErrorResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)

	assert.Equal(t, "Please, call fetch before.", errorResponse.Error)
}
