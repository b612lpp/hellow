package infra

import (
	"io"
	"log/slog"
)

func NewLogger(w io.Writer) *slog.Logger {
	l := slog.New(slog.NewTextHandler(w, nil))
	return l
}
