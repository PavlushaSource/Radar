package utils

import (
	"context"
	"time"
)

func WithTicker(ctx context.Context, ticker <-chan time.Time, action func()) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker:
				action()
			}
		}
	}()
}
