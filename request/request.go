// package request, handles request and getting infomration about commit frequency
package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// FormUrl creates url to send a request to api.github.
func FormUrl(username string) string {

	link := fmt.Sprintf("https://api.github.com/users/%v/events", username)
	return link
}

// Send, sends a get request to api.github, reads dates of
// all recent commits that api.github containts and returns them as []time.Time
func Send(apiLink string) ([]time.Time, error) {
	resp, err := http.Get(apiLink)
	if err != nil {
		return nil, fmt.Errorf("couldn't get a connection: %v", err)
	}
	defer resp.Body.Close()

	btResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	createdAts, err := unmarshalJson(btResp)
	if err != nil {
		return nil, err
	}

	return createdAts, nil
}

// unmarshalJson reads and returns all of dates of recent commits
func unmarshalJson(byteResp []byte) ([]time.Time, error) {
	var partialResp []CreatedAtResp
	err := json.Unmarshal(byteResp, &partialResp)
	if err != nil {
		return nil, err
	}

	var createdAts []time.Time
	for _, item := range partialResp {
		createdAts = append(createdAts, item.CreatedAt)
	}

	return createdAts, nil
}
