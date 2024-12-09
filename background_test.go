package lazy

import (
	"testing"
	"time"
)

func TestBackground(t *testing.T) {
	lazy := NewLazy(Background(&countingLoader{}))

	for i, exp := range []int{
		1, // sync load
		1, // return cached and trigger new load
		2, // return cached and trigger new load
		3, // return cached and trigger new load
	} {
		got, err := lazy.Get()
		if err != nil {
			t.Fatal(err)
		}

		if got != exp {
			t.Fatalf("call %d, expected %d, got %d", i+1, exp, got)
		}

		// make chance for the goroutine to load next element
		time.Sleep(time.Millisecond)
	}
}
