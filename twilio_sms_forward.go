package main

import (
	"net/http"

	"github.com/corybuecker/twilio-sms-forward/handlers"
	"github.com/corybuecker/twilio-sms-forward/logging"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	incomingMessageRoute := Route{URL: "/incoming_message", handler: handlers.CreateIncomingMessage, logger: new(logging.StandardOutLogger)}
	healthCheckRoute := Route{URL: "/", handler: handlers.CreateIncomingMessage, logger: new(logging.StandardOutLogger)}

	r.Handle(incomingMessageRoute.URL, incomingMessageRoute.HandlerWithLogging())
	r.Handle(healthCheckRoute.URL, healthCheckRoute.HandlerWithLogging())

	http.ListenAndServe(":8080", r)
}
