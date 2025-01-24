package stats

import (
	"time"
)

type StatResponse struct{
	Period time.Time `json:"period"`
	Sum int `json:"sum"`
}