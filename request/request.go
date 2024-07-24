package request

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
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

	// i want it to wait for couple of second if internet connection is trash
	// then drop it
	go func() {
		time.Sleep(time.Second * 5)
		if resp.StatusCode != http.StatusOK {
			fmt.Println("couldn't get a responsoe")
		}
	}()

	defer resp.Body.Close()

	strResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(strResp))
	return nil

}
