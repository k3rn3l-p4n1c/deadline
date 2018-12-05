package deadline

import (
	"sync"
	"context"
	"github.com/pkg/errors"
)

var timeoutExceededErr = errors.New("timeout exceeded")

func Run(ctx context.Context, function func(ctx context.Context)) error {
	return RunOr(ctx, function, func() {})
}

func RunOr(ctx context.Context, function func(ctx context.Context), afterTimeout func()) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	lock := sync.Mutex{}

	var finishSignal = make(chan struct{})
	defer close(finishSignal)

	var timeoutExceeded = false

	go func() {
		function(ctx)
		lock.Lock()
		if !timeoutExceeded {
			lock.Unlock()
			finishSignal <- struct{}{}
		}
	}()

	select {
	case <-finishSignal:
		lock.Lock()
		defer lock.Unlock()
		defer cancel()
		timeoutExceeded = true
		return nil

	case <-ctx.Done():
		lock.Lock()
		timeoutExceeded = true
		lock.Unlock()
		afterTimeout()
		return timeoutExceededErr
	}
}
