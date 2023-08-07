package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

func GetLogger(debug bool) zerolog.Logger {
	var writter io.Writer

	if debug {
		writter = zerolog.ConsoleWriter{Out: os.Stderr}
	} else {
		writter = os.Stderr
	}

	return zerolog.New(writter)
}
