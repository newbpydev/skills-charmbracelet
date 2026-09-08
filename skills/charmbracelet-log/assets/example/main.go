// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"charm.land/log/v2"
)

func newLogger(w io.Writer) *slog.Logger {
	handler := log.NewWithOptions(w, log.Options{ReportTimestamp: false})
	handler.SetFormatter(log.JSONFormatter)
	return slog.New(handler)
}
func main() {
	f, err := os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	logger := newLogger(f)
	logger.With("component", "example").Info("operation completed", "count", 3)
	fmt.Println("Wrote structured diagnostics to application.log")
}
