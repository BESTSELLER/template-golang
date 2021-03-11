package main

import (
	"github.com/{{cookiecutter.destination.git.owner + "/" + cookiecutter.destination.git.name}}/api"
	"github.com/{{cookiecutter.destination.git.owner + "/" + cookiecutter.destination.git.name}}/config"
	"github.com/{{cookiecutter.destination.git.owner + "/" + cookiecutter.destination.git.name}}/logger"
)

func main() {
	config.LoadEnvConfig()
	logger.Init()

	api.SetupRouter()
}
