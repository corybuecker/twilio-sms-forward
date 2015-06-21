package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type message struct {
	To          string `xml:"to,attr"`
	Description string `xml:",innerxml"`
}

type response struct {
	XMLName xml.Name `xml:"Response"`
	Message message  `xml:"Message"`
}

func CreateIncomingMessage(w http.ResponseWriter, r *http.Request) {
	var to []string
	to, error := r.URL.Query()["to"]
	if error != true {
		http.Error(w, "bad", 500)
		return
	}
	if len(to) != 1 {
		http.Error(w, "bad", 500)
		return
	}
	var responseData response
	responseData = response{}
	responseData.Message.To = to[0]
	responseData.Message.Description = "test"
	output, _ := xml.Marshal(responseData)

	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintln(w, string(output))
}
