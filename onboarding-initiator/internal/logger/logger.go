package logger

import (

	"log/slog"

	"os"

	"onboarding/internal/config"
)

func New(cfg config.Config) *slog.Logger {

	handler := slog.NewJSONHandler(

		os.Stdout,

		&slog.HandlerOptions{

			Level: slog.LevelInfo,
		},
	)

	return slog.New(handler)
}
