package lazy

import (
	"testing"
)

func TestOnce(t *testing.T) {
	lazy := NewLazy(Once(&countingLoader{}))

	got, err := lazy.Get()
	if err != nil {
		t.Fatal(err)
	}

	if got != 1 {
		t.Errorf("expected 1, got: %d", got)
	}

	got, err = lazy.Get()
	if err != nil {
		t.Fatal(err)
	}

	if got != 1 {
		t.Errorf("expected 1, got: %d", got)
	}
}
