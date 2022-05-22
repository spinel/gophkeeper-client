package logger

import (
	"os"

	"github.com/spinel/gophkeeper-client/pkg/config"

	"sync"

	"github.com/rs/zerolog"
)

// Logger is a logger :)
type Logger struct {
	*zerolog.Logger
}

var (
	logger Logger
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Logger {
	once.Do(func() {
		zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		cfg, _ := config.LoadConfig()
		// Set proper loglevel based on config
		switch cfg.LogLevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn", "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "err", "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel) // log info and above by default
		}
		logger = Logger{&zeroLogger}
	})
	return &logger
}
