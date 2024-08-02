package request

import (
	"fmt"
	"io"
	"net/http"
)

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

	fmt.Println(string(bytebody))
	return nil

}

func formUrl(uname string) string {
	url := fmt.Sprintf("https://api.github.com/users/%v/events", uname)
	return url

}
