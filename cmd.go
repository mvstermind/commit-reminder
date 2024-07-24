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
	flag.Parse()

	link := request.FormUrl(*username)
	err := request.Send(link)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

}
