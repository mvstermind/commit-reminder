package request

import "time"

type CreatedAtResp struct {
	CreatedAt time.Time `json:"created_at"`
}
