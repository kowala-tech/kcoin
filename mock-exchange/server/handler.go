package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/patrickmn/go-cache"
)

type FetchDataRequest struct {
	Sell []Value `json:"sell"`
	Buy  []Value `json:"buy"`
}

type Value struct {
	Amount float64
	Rate   float64
}

type FetchDataHandler struct {
	cache *cache.Cache
}

func (h *FetchDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	request := FetchDataRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error %s", err)
		return
	}

	h.cache.Set("last-request", request, cache.NoExpiration)
}
