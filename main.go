package main

import (
	"fmt"
	"log"
	"os"

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
		LongCmd:  "save",
		Usage:    "Save username",
		Required: false,
	}

	argList := minicli.AddArguments(&usernameArg, &styleArg)
	argMap := argList.Execute()

	username, ok := argMap["-u"]
	if !ok {
		return
	}

	todaysCommits, err := request.Get(username)
	if err != nil {
		log.Println("error: ", err)

	}

	save, ok := argMap["-s"]
	if ok {

		saveUsername(save)
	}

	printCommits(todaysCommits, username)

}

func printCommits(commits int, username string) {
	switch commits {

	case 0:
		fmt.Printf("%s haven't commited yet today\n", username)
	case 1:
		fmt.Printf("%s commited once today\n", username)

	case 2:
		fmt.Printf("%s commited twice today\n", username)

	default:
		fmt.Printf("%s commited %d today\n", username, commits)
	}
}

func saveUsername(uname string) bool {

	os.WriteFile("./username.txt", []byte(uname), 644)

}
