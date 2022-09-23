package zap_test

import (
	"time"
	"testing"
	"go.uber.org/zap"
	_ "go.uber.org/zap/zapcore"

	. "github.com/issenn/game-scraper/pkg/logger/adapter/zap"
)


func TestZapAdapter_New(t *testing.T) {
	New()
}

const url = "http://example.com"

func TestZapAdapter_Logger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestZapAdapter_SugaredLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func TestZapAdapter_convert(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("SugaredLogger failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("SugaredLogger Failed to fetch URL: %s", url)

	plain := sugar.Desugar()
	plain.Info("Logger failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestZapAdapter_LoggerDPanic(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger.DPanic("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
