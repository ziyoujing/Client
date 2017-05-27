package utils

import "time"

// Wait waits s seconds
func Wait(s time.Duration) {
	time.Sleep(s * time.Second)
}
