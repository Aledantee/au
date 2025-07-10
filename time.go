package au

import (
	"context"
	"time"
)

// SleepContext pauses execution for the specified duration or until the context is done, whichever occurs first.
func SleepContext(ctx context.Context, dur time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(dur):
	}
}
