package lazy

import "testing"

func TestFixed(t *testing.T) {
	exp := "test"
	lazy := NewLazy(Fixed(exp))

	got, err := lazy.Get()
	if err != nil {
		t.Fatal(err)
	}

	if got != exp {
		t.Errorf("expected %s, got %s", exp, got)
	}

	got, err = lazy.Get()
	if err != nil {
		t.Fatal(err)
	}

	if got != exp {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
