package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mvstermind/commit-remainder/request"
)

func ParseArgs() time.Time {

	if len(os.Args) <= 1 {
		fmt.Println("use -h or --help")
		return time.Time{}
	}
	username := flag.String("u", "", "github username")
	flag.Parse()

	link := request.FormUrl(*username)
	jsonTime, err := request.Send(link)

	if err != nil {
		fmt.Println("error: ", err)
		return time.Time{}
	}
	latestCommit := jsonTime[1]
	return latestCommit

}
