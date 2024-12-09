package lazy

import "time"

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (cf ClockFunc) Now() time.Time {
	return cf()
}

var DefaultClock = ClockFunc(time.Now)
