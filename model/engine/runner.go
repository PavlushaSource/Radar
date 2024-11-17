package engine

import (
	"context"
)

type Runner interface {
	Run(ctx context.Context) <-chan State
}

type runner struct {
	bufferSize int64
	engine     Engine
}

func CreateRunner(engine Engine, bufferSize int64) Runnable {
	return &Runner{
		engine:     engine,
		bufferSize: bufferSize,
	}
}

func (r *Runner) Run(ctx context.Context) <-chan State {
	resCh := make(chan State, r.bufferSize)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(resCh)
				return
			default:
				resCh <- r.engine.Run()
			}
		}
	}()

	return resCh
}
