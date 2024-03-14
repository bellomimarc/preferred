package logger

import (
	"context"
	"fmt"
	"os"
	"preferred/utils"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func setLoggingMode(mode string) {
	if mode == "" {
		mode = "pretty"
	}

	switch mode {
	case "pretty":
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:           os.Stderr,
			TimeFormat:    "2006-01-02 03:04:05PM",
			FormatMessage: formatMessage})
	case "json":
		log.Logger = log.Output(os.Stderr)
		zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999Z07:00"
	}
}

func formatMessage(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%s |", i)
}

func setLoggingLevel(level string) {
	if level == "" {
		level = "WARN"
	}

	level = strings.ToUpper(level)

	switch level {
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "ALL":
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
}

func InitLog() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	setLoggingLevel(os.Getenv("LOG_LEVEL"))
	setLoggingMode(os.Getenv("LOG_MODE"))
}

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()

	value := ctx.Value(utils.TraceId)
	if value != nil {
		e.Str("t_id", value.(string))
	}
}

var mylogger zerolog.Logger

func init() {
	mylogger = zerolog.New(os.Stdout)
	mylogger = mylogger.Hook(TracingHook{})
}

func WithContext(ctx context.Context) context.Context {
	return mylogger.WithContext(ctx)
}

func Error() *zerolog.Event {
	return mylogger.Error()
}

func Info() *zerolog.Event {
	return mylogger.Info()
}

func Debug() *zerolog.Event {
	return mylogger.Debug()
}

func Warn() *zerolog.Event {
	return mylogger.Warn()
}

func Trace() *zerolog.Event {
	return mylogger.Trace()
}

func Fatal() *zerolog.Event {
	return mylogger.Fatal()
}
