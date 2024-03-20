package benchmark

import (
	"time"
)

type TimeTracker struct {
}

func (timeTracker TimeTracker) Track(f func()) string {
	start := time.Now()
	f()
	return time.Since(start).String()
}
