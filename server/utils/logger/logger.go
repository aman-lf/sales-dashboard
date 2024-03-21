package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	once sync.Once
	log  zerolog.Logger
)

func InitLogger() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
			return file + ":" + strconv.Itoa(line)
		}

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.InfoLevel) // default to INFO
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Caller().
			Logger()
	})

	return log
}

func Debug(s string) {
	logger := InitLogger()
	logger.Debug().Msg(s)
}

func Info(s string) {
	logger := InitLogger()
	logger.Info().Msg(s)
}

func Warn(s string) {
	logger := InitLogger()
	logger.Warn().Msg(s)
}

func Error(s string, err error) {
	logger := InitLogger()
	logger.Error().Err(err).Msg(s)
}

func Fatal(s string) {
	logger := InitLogger()
	logger.Fatal().Msg(s)
}

func Panic(s string) {
	logger := InitLogger()
	logger.Panic().Msg(s)
}
