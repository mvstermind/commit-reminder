// package cmd parses command line arguments and sends request to api.github
package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mvstermind/commit-remainder/request"
)

// ParseArgs func reads sys args passed during program call.
// If proper flag for username is passed "-u" sends request to api.github
// returns latest commit date
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
	latestCommit := jsonTime[0]
	return latestCommit

}
