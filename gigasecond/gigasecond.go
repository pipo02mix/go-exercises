// Package gigasecond implements a method to add a 10^9 seconds to the provided time argument
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to the time passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1e9))
}
