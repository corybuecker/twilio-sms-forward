package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/corybuecker/twilio-sms-forward/logging"
)

type Route struct {
	URL     string
	handler func(http.ResponseWriter, *http.Request)
	logger  logging.Logger
}

func (route *Route) HandlerWithLogging() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		route.handler(w, r)
		route.logger.Log(fmt.Sprintf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start)))
	})
}
