package lazy

import (
	"context"
)

type Lazy[T any] struct {
	loader Loader[T]
}

func NewLazy[T any](loader Loader[T]) *Lazy[T] {
	return &Lazy[T]{
		loader: loader,
	}
}

func (l Lazy[T]) Get() (T, error) {
	return l.GetContext(context.Background())
}

func (l Lazy[T]) GetContext(ctx context.Context) (T, error) {
	return l.loader.Load(ctx)
}

type Loader[T any] interface {
	Load(ctx context.Context) (T, error)
}

type LoaderFunc[T any] func(ctx context.Context) (T, error)

func (f LoaderFunc[T]) Load(ctx context.Context) (T, error) {
	return f(ctx)
}
