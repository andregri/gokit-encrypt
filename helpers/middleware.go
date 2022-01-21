package helpers

import (
	"context"
	"time"

	log "github.com/go-kit/kit/log"
)

// LoggingMiddleware wraps the incoming requests with logs
type LogginMiddleware struct {
	Logger log.Logger
	Next   EncryptService
}

// Encrypt logs the encryption requests
func (mw LogginMiddleware) Encrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "encrypt",
			"key", key,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

// Decrypt logs the encryption requests
func (mw LogginMiddleware) Decrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "decrypt",
			"key", key,
			"message", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
