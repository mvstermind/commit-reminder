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
	if commitTime.Day() != time.Now().Day() {
		fmt.Println("you commited tday")
		return true
	}
	fmt.Println("you haven't commited tday")
	return false

}
