package main

import (
	"flag"
	"fmt"
	"os"
)

func ParseArgs() {

	if len(os.Args) <= 1 {
		fmt.Println("use -h or --help")
		return
	}
	username := flag.String("u", "", "github username")

	flag.Parse()
}
