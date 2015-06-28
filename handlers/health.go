package handlers

import "net/http"

// HealthCheck is a simple 200 response to test that the service is running
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
