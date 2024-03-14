package logger

import (
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

var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(os.Stdout)
	Logger = Logger.Hook(TracingHook{})
}
