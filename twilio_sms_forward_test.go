package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/corybuecker/twilio-sms-forward/handlers"
	"github.com/gorilla/mux"
)

type TestLogger struct {
	Logged bool
}

func (logger *TestLogger) Log(message string) {
	logger.Logged = true
}

var (
	server *httptest.Server

	incomingMessageRoute *Route
)

func init() {
	r := mux.NewRouter().StrictSlash(true)

	incomingMessageRoute = &Route{URL: "/incoming_message", handler: handlers.CreateIncomingMessage, logger: new(TestLogger)}

	r.Handle(incomingMessageRoute.URL, incomingMessageRoute.HandlerWithLogging())

	server = httptest.NewServer(r)
}

func TestResponseStatus(t *testing.T) {
	resp, _ := http.Get(server.URL + "/incoming_message?to=1")
	if resp.StatusCode != 200 {
		t.Errorf("expected a 200 status code")
	}
}
func TestErrorResponseStatus(t *testing.T) {
	resp, _ := http.Get(server.URL + "/incoming_message")
	if resp.StatusCode != 500 {
		t.Errorf("expected a 500 status code")
	}
}
func TestLogging(t *testing.T) {
	http.Get(server.URL + "/incoming_message?to=1")

	if incomingMessageRoute.logger.(*TestLogger).Logged != true {
		t.Errorf("expected logger to have been called")
	}
}

func TestToNumber(t *testing.T) {
	response, _ := http.Get(server.URL + "/incoming_message?to=12345")
	contents, _ := ioutil.ReadAll(response.Body)

	if strings.Contains(string(contents), "12345") != true {
		t.Errorf("body to contain to number")
	}
}
