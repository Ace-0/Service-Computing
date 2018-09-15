package service

import "net/http"

// StatusNotImplemented 501
const (
	StatusNotImplemented = 501
)

// NotImplemented a Handler
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", StatusNotImplemented)
}

// NotImplementedHandler return a Handler
func NotImplementedHandler() http.HandlerFunc {
	return http.HandlerFunc(NotImplemented)
}
