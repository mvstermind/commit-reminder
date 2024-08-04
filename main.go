package main

import (
	"log"

	"github.com/mvstermind/commit-reminder/prettify"
	"github.com/mvstermind/commit-reminder/request"
	minicli "github.com/mvstermind/mini-cli"
)

func main() {
	usernameArg := minicli.Arg{
		ShortCmd: "u",
		LongCmd:  "user",
		Usage:    "Add username flag to get info about commit frequency",
		Required: true,
	}

	styleArg := minicli.Arg{
		ShortCmd: "s",
		LongCmd:  "style",
		Usage:    "Modify output of displayed commit info",
		Required: false,
	}

	argList := minicli.AddArguments(&usernameArg, &styleArg)
	argMap := argList.Execute()

	username := argMap["-u"]
	style := argMap["-s"]

	todaysCommits, err := request.Get(username)
	if err != nil {
		log.Println("error: ", err)
	}
	prettify.Output(todaysCommits, style, username)

}
