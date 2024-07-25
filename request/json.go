package request

import "time"

// CreatedAtResp, hold information about commit time
type CreatedAtResp struct {
	CreatedAt time.Time `json:"created_at"`
}
