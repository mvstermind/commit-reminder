// entry point
package main

import (
	"fmt"
	"time"

	"github.com/mvstermind/commit-remainder/cmd"
)

func main() {
	cmtTime := cmd.ParseArgs()
	status := commitTdayStatus(cmtTime)

	fmt.Println(status)

}

// Validates checks if person commited tday returning true if commit was today
// and false if person hasn't commited yet
func commitTdayStatus(commitTime time.Time) bool {
	commitTimeDay := commitTime.Day()
	currDay := time.Now().Day()

	if commitTimeDay == currDay {
		fmt.Println("You have commited tday!!")
		return true
	}

	fmt.Println("You haven't commited tday :(")
	return false

}
