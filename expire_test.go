package lazy

import (
	"testing"
	"time"
)

func TestExpire(t *testing.T) {
	clock := MockedClock(time.Now(),
		0*time.Second, // obtaining lock
		0*time.Second, // register cache item
		time.Second,   // 2nd call
		time.Second,   // 3rd call
		time.Second,   // 4th call
	)
	lazy := NewLazy(ExpireClock(2*time.Second, clock, &countingLoader{}))

	for i, exp := range []int{
		1,
		1,
		1,
		2,
		2,
		2,
		3,
	} {
		got, err := lazy.Get()
		if err != nil {
			t.Fatal(err)
		}

		if got != exp {
			t.Fatalf("call %d: expected %d, got %d", i+1, exp, got)
		}
	}
}
