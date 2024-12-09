package lazy

import (
	"context"
	"sync"
	"time"
)

type expire[T any] struct {
	sync.Mutex
	clock        Clock
	lastLoadedAt time.Time
	after        time.Duration
	loaded       bool
	cached       T
	loader       Loader[T]
}

func Expire[T any](after time.Duration, loader Loader[T]) *expire[T] {
	return ExpireClock(after, DefaultClock, loader)
}

func ExpireClock[T any](after time.Duration, clock Clock, loader Loader[T]) *expire[T] {
	return &expire[T]{
		after:  after,
		clock:  clock,
		loader: loader,
	}
}

func (r *expire[T]) Load(ctx context.Context) (T, error) {
	now := r.clock.Now()
	expired := r.lastLoadedAt.Add(r.after).Before(now)
	if !expired && r.loaded {
		return r.cached, nil
	}

	r.Lock()
	defer r.Unlock()

	now = r.clock.Now()
	expired = r.lastLoadedAt.Add(r.after).Before(now)
	if !expired && r.loaded {
		return r.cached, nil
	}

	result, err := r.loader.Load(ctx)
	if err != nil {
		return *new(T), err
	}

	r.cached = result
	r.loaded = true
	r.lastLoadedAt = r.clock.Now()

	return result, nil
}
