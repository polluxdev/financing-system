package helper

import (
	"time"
)

func GenerateTimeDuration(n int, d time.Duration) time.Duration {
	return time.Duration(n) * d
}
