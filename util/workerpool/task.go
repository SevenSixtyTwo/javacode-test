package workerpool

import (
	"context"
	"errors"
	"time"
)

type Task struct {
	Err chan error
	// Res chan interface{}

	f func() error

	timeout time.Duration
}

var (
	defaultDuration = time.Second * 20
	errTimeoutTask  = errors.New("error timeout task from worker")
)

func NewTask(f func() error, timeout *time.Duration) *Task {
	if timeout == nil {
		timeout = &defaultDuration
	}

	return &Task{
		f:       f,
		timeout: *timeout,
		// Res:     make(chan interface{}, 2),
		Err: make(chan error, 2),
	}
}

type response struct {
	// Res	interface{}
	Err error
}

func process(t *Task) {
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	responseCh := make(chan response)

	go func() {
		err := t.f()
		responseCh <- response{err}
	}()

	select {
	case <-ctx.Done():
		t.Err <- errTimeoutTask

	case res := <-responseCh:
		t.Err <- res.Err
	}

	// close(t.Err)
}
