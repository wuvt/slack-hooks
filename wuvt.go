package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nlopes/slack"
)

const trackmanURL = "https://trackman-fm.apps.wuvt.vt.edu/api"
const djLinkURL = "https://www.wuvt.vt.edu/playlists/dj/%d"
const trackLinkURL = "https://www.wuvt.vt.edu/playlists/track/%d"

type TrackmanLatestTrackResponse struct {
	Album     string `json:"album"`
	Artist    string `json:"artist"`
	DJ        string `json:"dj"`
	DJID      int    `json:"dj_id"`
	Label     string `json:"label"`
	Title     string `json:"title"`
	Listeners int    `json:"listeners"`
	TrackID   int    `json:"id"`
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

	trackLink := fmt.Sprintf("<%s|%s>", fmt.Sprintf(trackLinkURL, track.TrackID), track.Title)

	message := slack.NewBlockMessage(
		slack.NewSectionBlock(
			slack.NewTextBlockObject(
				slack.MarkdownType,
				fmt.Sprintf("*%s - %s*", track.Artist, trackLink),
				false,
				false,
			),
			[]*slack.TextBlockObject{
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("Album: %s", track.Album),
					false,
					false,
				),
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("Label: %s", track.Label),
					false,
					false,
				),
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("DJ: %s", djLink),
					false,
					false,
				),
				slack.NewTextBlockObject(
					slack.MarkdownType,
					fmt.Sprintf("Listeners: %d", track.Listeners),
					false,
					false,
				),
			},
			nil,
		),
	)

	message.Msg.ResponseType = "in_channel"

	output, err := json.Marshal(message)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
