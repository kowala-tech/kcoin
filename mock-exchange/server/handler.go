package server

import (
	"encoding/json"
	"net/http"

	"github.com/kowala-tech/kcoin/mock-exchange/app"
	"github.com/kowala-tech/kcoin/mock-exchange/exchange"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type FetchDataHandler struct {
	cache *cache.Cache
}

func (h *FetchDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithFields(logrus.Fields{
		"call":   "fetch",
		"method": r.Method,
	})

	if r.Method != http.MethodPost {
		logger.Warn("invalid method called")

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	request := app.Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.WithError(err).WithField("request", request).Warn("error creating request")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.cache.Set(app.CacheRequestKey, request, cache.NoExpiration)
	logger.WithField("request", request).Info("saved request in cache")
}

type GetRatesHandler struct {
	cache       *cache.Cache
	transformer exchange.Transformer
}

func (h *GetRatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithFields(logrus.Fields{
		"call":   "getRate",
		"method": r.Method,
	})

	if r.Method != http.MethodGet {
		logger.Warn("invalid method called")

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req, ok := h.cache.Get(app.CacheRequestKey)
	if !ok {
		jsonErrResp, err := json.Marshal(ErrorResponse{Error: "Please, call fetch before."})
		if err != nil {
			logger.WithError(err).Warn("error creating error response")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErrResp)

		logger.Info("msg", "get rates api method was called before fetching mocked info")
		return
	}

	tResp, err := h.transformer.Transform(req.(app.Request))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonErrResp, errJson := json.Marshal(ErrorResponse{Error: "There was a problem getting data from request."})
		if errJson != nil {
			logger.WithError(err).Warn("error creating error response")
			return
		}

		w.Write(jsonErrResp)
		logger.WithError(err).Warn("msg", "there was a problem when trying to decodify data from cache")
		return
	}

	w.Write([]byte(tResp))
	logger.WithField("response", tResp).Info("msg", "request for data rates accomplished")
}
