package prettify

import (
	"fmt"
	"strconv"
)

// todo finish this shit
func Output(value int, style, username string) {

	valueStr := strconv.Itoa(value)

	switch valueStr {
	case "1":
		valueStr = "once today"
	case "2":
		valueStr = "twice today"

	default:
		valueStr = fmt.Sprintf("%v times today", valueStr)

	}

	switch style {
	case "d-_-b":
		fmt.Printf("d-_-b %s commited today: %v times d-_-b\n", username, valueStr)

	case "":
		fmt.Printf("%v commited: %v\n", username, valueStr)

	}
}
