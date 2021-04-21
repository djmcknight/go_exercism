// Package gigasecond calculates the time after you've added 1e9 to a time.
package gigasecond

import (
	"time"
)

// AddGigasecond adds 1e9 seconds to the variable time passed to the function, returning Time.time in the future.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second* 1e9 )
}
