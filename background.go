package lazy

import (
	"context"
	"sync"
)

type background[T any] struct {
	sync.Mutex
	loaded bool
	cached T
	err    error
	loader Loader[T]
}

func Background[T any](loader Loader[T]) *background[T] {
	return &background[T]{
		loader: loader,
	}
}

func (b *background[T]) Load(ctx context.Context) (T, error) {
	if b.loaded {
		go b.load(context.Background())
		return b.cached, b.err
	}

	return b.load(ctx)
}

func (b *background[T]) load(ctx context.Context) (T, error) {
	b.Lock()
	defer b.Unlock()

	result, err := b.loader.Load(ctx)
	if err != nil {
		b.cached = *new(T)
		b.err = err
		b.loaded = true
		return b.cached, b.err
	}

	b.cached = result
	b.err = nil
	b.loaded = true

	return result, nil
}
