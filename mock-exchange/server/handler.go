package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kowala-tech/kcoin/mock-exchange/app"

	"github.com/kowala-tech/kcoin/mock-exchange/exchange"

	"github.com/patrickmn/go-cache"
)

type ErrorResponse struct {
	Error string `json:"error"`
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

	request := app.Request{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error %s", err)
		return
	}

	h.cache.Set(app.CacheRequestKey, request, cache.NoExpiration)
}

type GetRatesHandler struct {
	cache       *cache.Cache
	transformer exchange.Transformer
}

func (h *GetRatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req, ok := h.cache.Get(app.CacheRequestKey)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		jsonErrResp, err := json.Marshal(ErrorResponse{Error: "Please, call fetch before."})
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		w.Write(jsonErrResp)
		return
	}

	tResp, err := h.transformer.Transform(req.(app.Request))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonErrResp, err := json.Marshal(ErrorResponse{Error: "Please, call fetch before."})
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		w.Write(jsonErrResp)
		return
	}

	w.Write([]byte(tResp))
}
