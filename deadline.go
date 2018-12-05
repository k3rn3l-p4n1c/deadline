package deadline

import (
	"sync"
	"context"
)

func Run(ctx context.Context, function func(ctx context.Context)) {
	RunOr(ctx, function, func() {})
}

func RunOr(ctx context.Context, function func(ctx context.Context), afterTimeout func()) {
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
	case <-ctx.Done():
		timeoutExceeded = true
		afterTimeout()
	}
	lock.Lock()
	cancel()
	timeoutExceeded = true
	lock.Unlock()
}
