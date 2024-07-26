// package cmd parses command line arguments and sends request to api.github
package cmd

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/mvstermind/commit-remainder/request"
)

// ParseArgs func reads sys args passed during program call.
// If proper flag for username is passed "-u" sends request to api.github
// returns latest commit date
func ParseArgs() (time.Time, error) {

	if len(os.Args) <= 1 {
		return time.Time{}, errors.New("use -h or --help")
	}
	username := flag.String("u", "", "github username")
	flag.Parse()

	link := request.FormUrl(*username)
	jsonTime, err := request.Send(link)

	if err != nil {
		return time.Time{}, err
	}
	latestCommit := jsonTime[0]
	return latestCommit, nil

}
