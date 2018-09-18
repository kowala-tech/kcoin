package server

import (
	"github.com/gorilla/mux"
	"github.com/kowala-tech/kcoin/mock-exchange/exchange/exrates"
	"github.com/patrickmn/go-cache"
	"time"
)

func GetRouter() *mux.Router {
	c := cache.New(5*time.Minute, 10*time.Minute)

	r := mux.NewRouter()
	r.Handle("/api/fetch", &FetchDataHandler{cache:c})
	r.Handle("/api/exrates/get", &GetRatesHandler{cache:c, transformer: &exrates.Transformer{}})

	return r
}
