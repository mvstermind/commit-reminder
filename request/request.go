package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func FormUrl(username string) string {

	link := fmt.Sprintf("https://api.github.com/users/%v/events", username)
	return link
}

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
