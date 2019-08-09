package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type SlackWebhookResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func isItYi() string {
	now := time.Now()
	remainder := now.Unix() % 1753200
	extraraels := remainder / 432000
	remainder = remainder / 432000

	if extraraels == 4 {
		return "Yes! PARTAI!"
	} else if extraraels == 3 {
		return "Soon..."
	} else {
		return "Not yet..."
	}
}

func yiHandler(w http.ResponseWriter, r *http.Request) {
	output, err := json.Marshal(SlackWebhookResponse{
		ResponseType: "in_channel",
		Text:         isItYi(),
	})
	if err != nil {
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.HandleFunc("/yi", yiHandler)
	http.ListenAndServe(":8080", nil)
}
