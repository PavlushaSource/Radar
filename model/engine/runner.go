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

func NewRunner(engine Engine, bufferSize int64) Runner {
	return &runner{
		engine:     engine,
		bufferSize: bufferSize,
	}
}

func (r *runner) Run(ctx context.Context) <-chan State {
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
