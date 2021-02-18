package logger

import (
	"strings"

	"github.com/{{cookiecutter.destination.git.owner + "/" + cookiecutter.destination.git.name}}/config"
	"github.com/rs/zerolog"
)

// Init sets the global loglevel
func Init() {
	// default is info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if config.EnvVars.LogLevel == "" {
		return
	}
	if strings.ToLower(config.EnvVars.LogLevel) == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
