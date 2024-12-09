package lazy

import (
	"fmt"
	"time"
)

type mockedClock struct {
	next      time.Time
	at        int
	durations []time.Duration
}

func MockedClock(start time.Time, durations ...time.Duration) *mockedClock {
	return &mockedClock{
		next:      start,
		at:        0,
		durations: durations,
	}
}

func (m *mockedClock) Now() time.Time {
	result := m.next

	if len(m.durations) > 0 {
		m.next = m.next.Add(m.durations[m.at%len(m.durations)])
		m.at++
	}

	return result
}

func (m mockedClock) String() string {
	return fmt.Sprintf("mocked clock time: %s, at tick %d", m.next.Format(time.DateTime), m.at)
}
