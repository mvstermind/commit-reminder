package main

import (
	"errors"
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
		Required: false,
	}

	styleArg := minicli.Arg{
		ShortCmd: "s",
		LongCmd:  "save",
		Usage:    "Save username",
		Required: false,
	}

	argList := minicli.AddArguments(&usernameArg, &styleArg)
	argMap := argList.Execute()

	username := argMap["-u"]

	if username == "" {
		var err error
		username, err = loadUname()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	todaysCommits, err := request.Get(username)
	if err != nil {
		log.Println("error: ", err)

	}

	save := argMap["-s"]
	if save == "" {
		saveUsername(username)
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

func saveUsername(uname string) {

	dir, err := os.Getwd()
	if err != nil {
		return
	}

	fileNdir := fmt.Sprintf("%v/username.txt", dir)

	file, err := os.OpenFile(fileNdir, os.O_WRONLY, 0666)
	if err != nil {
		err := fmt.Errorf("Error: %v\n", err)
		fmt.Printf(err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(uname)
	if err != nil {
		fmt.Println(err)
	}

}

func loadUname() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fileNdir := fmt.Sprintf("%v/username.txt", dir)
	content, err := os.ReadFile(fileNdir)

	if len(content) == 0 {
		return "", errors.New("no username saved")
	}

	if err != nil {
		return "", err
	}

	fileContent := string(content)
	return fileContent, nil

}
