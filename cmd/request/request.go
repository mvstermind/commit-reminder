package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
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

	// Fill in json response into slice
	var events []Event
	err = json.Unmarshal(bytebody, &events)
	if err != nil {
		return err
	}

	// Modify json response to match format of currentDate (declared in howManyFunc)
	var refactoredDate []string
	for _, value := range events {
		valid := strings.Split(value.CreatedAt, "T")
		refactoredDate = append(refactoredDate, valid[0])

	}
	todayCommmits := commitsToday(refactoredDate)

	fmt.Println(todayCommmits)
	return nil
}

func formUrl(uname string) string {
	url := fmt.Sprintf("https://api.github.com/users/%v/events", uname)
	return url

}

func commitsToday(commits []string) int {
	currentDate := time.Now().Format("2006-01-02")

	var commitsToday int

	for _, commit := range commits {
		if commit == currentDate {
			commitsToday++
		}
	}

	return commitsToday

}
