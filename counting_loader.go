package lazy

import "context"

type countingLoader struct {
	count int
}

func (cl *countingLoader) Load(ctx context.Context) (int, error) {
	cl.count++
	return cl.count, nil
}
