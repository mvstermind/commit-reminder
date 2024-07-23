package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mvstermind/commit-remainder/request"
)

func ParseArgs() {

	if len(os.Args) <= 1 {
		fmt.Println("use -h or --help")
		return
	}
	username := flag.String("u", "", "github username")

	link := request.FormUrl(username)
	request.Send(link)

	flag.Parse()
}
