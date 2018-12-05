package deadline

import (
	"testing"
	"context"
	"time"
)

func TestProperRun(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	i := 0
	Run(ctx, func(_ context.Context) {
		i++
	})
	if i != 1 {
		t.Fatal("function have not been executed properly")
	}
}

func TestTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	i := 0
	startTime := time.Now()
	Run(ctx, func(_ context.Context) {
		time.Sleep(3 * time.Second)
		i++
	})
	took := time.Now().Sub(startTime)
	if i != 0 {
		t.Fatal("function have not been timeouted")
	}
	if int(took.Seconds()) > 1 {
		t.Fatal("function take time more than timeout")
	}
}

func TestContextPropagation(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	i := 0
	Run(ctx, func(innerCtx context.Context) {
		for {
			time.Sleep(1 * time.Millisecond)
			i++
			if innerCtx.Err() != nil {
				break
			}
		}
	})
	time.Sleep(1 * time.Second)

	if i > int(time.Second/time.Millisecond) {
		println(i)
		t.Fatal("context didn't propagate timeout")
	}
}

func TestRunOrTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	afterTimeoutIsExecuted := false
	RunOr(ctx, func(_ context.Context) {
		time.Sleep(2*time.Second)
	}, func() {
		afterTimeoutIsExecuted = true
	})
	time.Sleep(1 * time.Second)

	if !afterTimeoutIsExecuted {
		t.Fatal("after timeout function is not executed")
	}
}


func TestRunOr(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	afterTimeoutIsExecuted := false
	RunOr(ctx, func(_ context.Context) {
		time.Sleep(1*time.Second)
	}, func() {
		afterTimeoutIsExecuted = true
	})
	time.Sleep(2 * time.Second)

	if afterTimeoutIsExecuted {
		t.Fatal("after timeout function is executed")
	}
}
