package log

import (
	"os"

	"github.com/op/go-logging"
)

func levelLogger() *logging.Logger {
	var logger *logging.Logger = logging.MustGetLogger("example")
	var format logging.Formatter = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶▶▶ %{level:.4s} %{id:03x}%{color:reset} %{message} ▶▶▶`,
	)
	var backend *logging.LogBackend = logging.NewLogBackend(os.Stderr, "", 0)
	var backendFormatter logging.Backend = logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	return logger

}

