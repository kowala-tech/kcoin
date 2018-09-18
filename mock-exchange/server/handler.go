package server

import (
	"encoding/json"
	"fmt"
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
	if r.Method != http.MethodPost {
		logrus.WithFields(logrus.Fields{
			"call":   "fetch",
			"method": r.Method,
		}).Warn("invalid method called")

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	request := app.Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"call":    "fetch",
			"request": request,
		}).Warn(fmt.Sprintf("error creating request: %s", err))

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.cache.Set(app.CacheRequestKey, request, cache.NoExpiration)
	logrus.WithFields(logrus.Fields{
		"call":    "fetch",
		"request": request,
	}).Info("saved request in cache")
}

type GetRatesHandler struct {
	cache       *cache.Cache
	transformer exchange.Transformer
}

func (h *GetRatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logrus.WithFields(logrus.Fields{
			"call":   "getRate",
			"method": r.Method,
		}).Warn("invalid method called")

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req, ok := h.cache.Get(app.CacheRequestKey)
	if !ok {

		jsonErrResp, err := json.Marshal(ErrorResponse{Error: "Please, call fetch before."})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"call":  "getRate",
				"error": err,
			}).Warn("error creating error response")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErrResp)

		logrus.WithFields(logrus.Fields{
			"call": "getRate",
		}).Info("msg", "get rates api method was called before fetching mocked info")
		return
	}

	tResp, err := h.transformer.Transform(req.(app.Request))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonErrResp, err := json.Marshal(ErrorResponse{Error: "There was a problem getting data from request."})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"call":  "getRate",
				"error": err,
			}).Warn("error creating error response")
			return
		}

		w.Write(jsonErrResp)
		logrus.WithFields(logrus.Fields{
			"call":  "getRate",
			"error": err,
		}).Warn("msg", "there was a problem when trying to decodify data from cache")
		return
	}

	w.Write([]byte(tResp))
	logrus.WithFields(logrus.Fields{
		"call": "getRate",
		"data": tResp,
	}).Info("msg", "request for data rates accomplished")
}
