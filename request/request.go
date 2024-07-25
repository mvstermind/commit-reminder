package request

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func FormUrl(username string) string {

	fmt.Println("value of username :", username)
	link := fmt.Sprintf("https://api.github.com/users/%v/events", username)
	return link
}

func Send(apiLink string) error {

	resp, err := http.Get(apiLink)
	if err != nil {
		return fmt.Errorf("couldn't get a connection %v", err)

	}

	defer resp.Body.Close()

	strResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(strResp))
	return nil

}
