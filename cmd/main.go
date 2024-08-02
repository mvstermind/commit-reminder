package main

import (
	"log"

	"github.com/mvstermind/did-i-commit/cmd/request"
	minicli "github.com/mvstermind/mini-cli"
)

func main() {
	userArg := minicli.Arg{
		ShortCmd: "u",
		LongCmd:  "user",
		Usage:    "Add username flag to get info about commit frequency",
		Required: true,
	}

	argList := minicli.AddArguments(&userArg)
	argMap := argList.Execute()
	err := request.Get(argMap["-u"])
	if err != nil {
		log.Println("error: ", err)
	}

}
