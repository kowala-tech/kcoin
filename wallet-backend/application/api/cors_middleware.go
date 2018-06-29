package api

import "net/http"

//setHandlerCors is a wrapper function to convert a normal handler to a
//allow all access control.
func setHandlerCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
