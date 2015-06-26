package handlers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

type message struct {
	To   string `xml:"to,attr"`
	Body string `xml:",innerxml"`
}

type response struct {
	XMLName xml.Name `xml:"Response"`
	Message message  `xml:"Message"`
}

func parseQueryParams(request *http.Request) (map[string]string, error) {
	m := make(map[string]string)
	m["to"] = request.FormValue("to")
	m["from"] = request.FormValue("From")
	m["body"] = request.FormValue("Body")
	if len(m["to"]) == 0 || len(m["from"]) == 0 || len(m["body"]) == 0 {
		return m, errors.New("missing values")
	}
	return m, nil
}

// CreateIncomingMessage is the handler to generate a forwarding XML response
func CreateIncomingMessage(w http.ResponseWriter, r *http.Request) {
	var (
		parameters   map[string]string
		err          error
		responseData response
	)

	if parameters, err = parseQueryParams(r); err != nil {
		http.Error(w, "", 500)
		return
	}
	responseData = response{}
	responseData.Message.To = parameters["to"]

	smsBody := fmt.Sprintf("%s: %s", parameters["from"], parameters["body"])
	if len(smsBody) > 160 {
		smsBody = smsBody[:160]
	}

	responseData.Message.Body = smsBody
	output, _ := xml.Marshal(responseData)

	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintln(w, string(output))
}
