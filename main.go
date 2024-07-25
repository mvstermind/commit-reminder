package main

import (
	"fmt"
	"time"
)

func main() {
	time := ParseArgs()
	fmt.Println(time)
	status := commitStatus(time)
	fmt.Println(status)

}

func commitStatus(committime time.Time) bool {

	if time.Since(committime) >= time.Hour {
		fmt.Println("commited more than hour ago")
		return true
	}

	return false
}
