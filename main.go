package main

import (
	"fmt"
	"time"
)

func main() {
	cmtTime := ParseArgs()
	status := commitTdayStatus(cmtTime)

	fmt.Println(status)

}

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
