package main

import (
	"net/http"
)

type SlackWebhookResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func main() {
	http.HandleFunc("/wuvt", wuvtHandler)
	http.HandleFunc("/yi", yiHandler)
	http.ListenAndServe(":8080", nil)
}
