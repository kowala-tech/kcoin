package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FetchDataRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type FetchDataHandler struct {
}

func (*FetchDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("%v", request)
}
