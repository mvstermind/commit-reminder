package main

import (
	"fmt"

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

	fmt.Printf("%v\n", argMap["-u"])
}
