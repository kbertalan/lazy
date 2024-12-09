package lazy

import (
	"context"
	"sync"
)

type once[T any] struct {
	sync.Mutex
	loaded bool
	cached T
	loader Loader[T]
}

func Once[T any](loader Loader[T]) *once[T] {
	return &once[T]{
		loader: loader,
	}
}

func (o *once[T]) Load(ctx context.Context) (T, error) {
	if o.loaded {
		return o.cached, nil
	}

	o.Lock()
	defer o.Unlock()

	if o.loaded {
		return o.cached, nil
	}

	result, err := o.loader.Load(ctx)
	if err != nil {
		return *new(T), err
	}

	o.cached = result
	o.loaded = true

	return result, nil
}
