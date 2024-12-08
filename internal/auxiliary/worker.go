package auxiliary

import (
	"context"
	"sync"
)

var (
	mutex = &sync.Mutex{}
)

type Context struct {
	PushError func(err error)
	context.Context
}

func NewWorker[A interface{}, R interface{}](ctx context.Context, slice []A, callback func(Context, A) R) ([]R, error) {

	if len(slice) == 0 {
		return nil, nil
	}

	newSlice := make([]R, len(slice))

	done := make(chan error, 1)
	count := make(chan int, 1)

	pusher := func(err error) {
		done <- err
	}

	aCtx := Context{
		PushError: pusher,
		Context:   ctx,
	}

	count <- 1

	for v, i := range slice {
		go goroutineMaker(i, newSlice, v, count, done, callback, aCtx)
	}

	err := <-done

	if err != nil {
		return nil, err
	}

	return newSlice, nil
}

func goroutineMaker[A interface{}, R interface{}](
	v A,
	newArr []R,
	i int,
	count chan int,
	done chan<- error,
	callback func(Context, A) R,
	ctx Context,
) {
	newArr[i] = callback(ctx, v)

	mutex.Lock()

	c := <-count
	// println(c)

	count <- c + 1

	mutex.Unlock()

	if c == len(newArr) {
		done <- nil
	}
}
