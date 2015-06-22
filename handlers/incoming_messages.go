package handlers

import (
	"encoding/xml"
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

func CreateIncomingMessage(w http.ResponseWriter, r *http.Request) {
	var to []string
	var from, body string

	to, _ = r.URL.Query()["to"]

	from = r.FormValue("From")
	body = r.FormValue("Body")

	if len(to) != 1 || len(from) == 0 || len(body) == 0 {
		http.Error(w, "Missing Query Parameters", 500)
		return
	}

	var responseData response
	responseData = response{}
	responseData.Message.To = to[0]

	smsBody := fmt.Sprintf("%s: %s", from, body)
	if len(smsBody) > 160 {
		smsBody = smsBody[:160]
	}

	responseData.Message.Body = smsBody
	output, _ := xml.Marshal(responseData)

	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintln(w, string(output))
}
