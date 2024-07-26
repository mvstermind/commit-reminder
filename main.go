// entry point
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mvstermind/commit-remainder/cmd"
)

func main() {
	cmtTime, err := cmd.ParseArgs()
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	commitTdayStatus(cmtTime)

}

// Validates checks if person commited tday returning true if commit was today
// and false if person hasn't commited yet
func commitTdayStatus(commitTime []time.Time) {
	commitTimeDay := commitTime[0].Day()
	currDay := time.Now().Day()

	for _, v := range commitTime {
		fmt.Println(v)

	}

	emoji := '\U0001F389'
	if commitTimeDay == currDay && len(commitTime) == 1 {
		fmt.Printf("\n%c  You have commited tday !! %c\n\n", emoji, emoji)
		return
	}

	if commitTimeDay == currDay && len(commitTime) > 1 {
		fmt.Printf("\n%c  You have commited tday: %v times!! %c\n\n", emoji, len(commitTime), emoji)
		return

	}

	fmt.Println("You haven't commited tday :(")

}
