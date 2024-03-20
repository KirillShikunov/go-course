package benchmark

import (
	"time"
)

type TimeTracker struct {
}

func (timeTracker TimeTracker) Track(f func()) int64 {
	start := time.Now()
	f()
	return time.Since(start).Nanoseconds()
}
