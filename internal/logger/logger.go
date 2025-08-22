package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)


type Logger = zerolog.Logger


func New(level string) Logger {
lvl, err := zerolog.ParseLevel(strings.ToLower(level))
if err != nil { lvl = zerolog.InfoLevel }
zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
l := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(lvl)
return l
}