package clock

import "time"

var fakeTime time.Time

// Now Now
func Now() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}

// Set Set
func Set(t time.Time) {
	fakeTime = t
}

// Reset Reset
func Reset() {
	fakeTime = time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
}
