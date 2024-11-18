package runner

import (
	"context"
	"github.com/PavlushaSource/Radar/model/engine"
)

type Runner interface {
	Run(ctx context.Context) <-chan engine.State
}

type runner struct {
	bufferSize int64
	engine     engine.Engine
}

func NewRunner(engine engine.Engine, bufferSize int64) Runner {
	return &runner{
		engine:     engine,
		bufferSize: bufferSize,
	}
}

func (r *runner) Run(ctx context.Context) <-chan engine.State {
	resCh := make(chan engine.State, r.bufferSize)

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
