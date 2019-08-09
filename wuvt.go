package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const trackmanURL = "https://trackman-fm.apps.wuvt.vt.edu/api"
const djLinkURL = "https://www.wuvt.vt.edu/playlists/dj/%d"

type TrackmanLatestTrackResponse struct {
	Album  string `json:"album"`
	Artist string `json:"artist"`
	DJ     string `json:"dj"`
	DJID   int    `json:"dj_id"`
	Label  string `json:"label"`
	Title  string `json:"title"`
}

func wuvtHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(trackmanURL + "/playlists/latest_track")
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to load latest track information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to load latest track information", http.StatusInternalServerError)
		return
	}

	var track TrackmanLatestTrackResponse
	err = json.Unmarshal(body, &track)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to load latest track information", http.StatusInternalServerError)
		return
	}

	var djLink string
	if track.DJID > 0 {
		djLink = fmt.Sprintf("<%s|%s>", fmt.Sprintf(djLinkURL, track.DJID), track.DJ)
	} else {
		djLink = track.DJ
	}

	output, err := json.Marshal(SlackWebhookResponse{
		ResponseType: "in_channel",
		Text:         fmt.Sprintf("*%s - %s*\nDJ: %s", track.Artist, track.Title, djLink),
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
