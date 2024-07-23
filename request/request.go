package request

import (
	"fmt"
)

func FormUrl(username *string) string {

	link := fmt.Sprintf("https://%v", username)
	return link
}

func Send(accountLink string) {

}
