package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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
		log.Print(err)
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
