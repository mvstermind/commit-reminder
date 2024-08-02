package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Event struct {
	CreatedAt string `json:"created_at"`
}

func Get(username string) error {
	link := formUrl(username)

	req, err := http.Get(link)
	if err != nil {
		return err
	}

	bytebody, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var events []Event

	err = json.Unmarshal(bytebody, &events)
	if err != nil {
		return err
	}

	for _, event := range events {
		fmt.Println(event)
	}

	return nil

}

func formUrl(uname string) string {
	url := fmt.Sprintf("https://api.github.com/users/%v/events", uname)
	return url

}
