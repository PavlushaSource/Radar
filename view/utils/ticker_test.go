package utils

import (
	"context"
	"testing"
	"time"
)

func TestWithTicker(t *testing.T) {
	ctx := context.Background()
	interval := 1 * time.Second

	timePrev := time.Now()
	action := func() {
		now := time.Now()
		if now.Sub(timePrev).Seconds()+(25*time.Millisecond).Seconds() >= interval.Seconds() &&
			now.Sub(timePrev).Seconds()-(25*time.Millisecond).Seconds() <= interval.Seconds() {
			// all good
		} else {
			t.Errorf("Ticker works incorrect, time in seconds for one action - (%v)", now.Sub(timePrev).Seconds())
		}
		timePrev = now
	}
	ticker := time.Tick(interval)

	WithTicker(ctx, ticker, action)
	time.Sleep(6 * time.Second)
}
