package lazy

import (
	"testing"
	"time"
)

func TestLazy(t *testing.T) {
	clock := MockedClock(time.Now(),
		0*time.Second,
		0*time.Second,
		time.Second,
		time.Second,
		time.Second,
	)
	lazy := NewLazy(Background(ExpireClock(2*time.Second, clock, &countingLoader{})))

	for i, exp := range []int{
		1, // sync load
		1, // return cached
		1, // return cached
		1, // return cached, background load
		2, // return cached
		2, // return cached
		2, // return cached, background load
		3, // return cached
		3, // return cached
	} {
		got, err := lazy.Get()
		if err != nil {
			t.Fatal(err)
		}

		if got != exp {
			t.Fatalf("call %d, expected %d, got: %d", i+1, exp, got)
		}

		time.Sleep(time.Millisecond)
	}
}
