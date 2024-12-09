package lazy

import "context"

type fixed[T any] struct {
	cached T
}

func Fixed[T any](value T) *fixed[T] {
	return &fixed[T]{
		cached: value,
	}
}

func (f *fixed[T]) Load(ctx context.Context) (T, error) {
	return f.cached, nil
}
