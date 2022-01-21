package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentationMiddleware measures 2 parameters: request count, latency
type InstrumentationMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           EncryptService
}

// Encrypt observes two variables for encrypt requests
func (mw InstrumentationMiddleware) Encrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{
			"method", "encrypt",
			"error", fmt.Sprint(err != nil),
		}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

// Decrypt observes two variables for decrypt requests
func (mw InstrumentationMiddleware) Decrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{
			"method", "decrypt",
			"error", "false",
		}

		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
