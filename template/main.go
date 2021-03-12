package main

import (
	"github.com/BESTSELLER/{{cookiecutter.name}}/api"
	"github.com/BESTSELLER/{{cookiecutter.name}}/config"
	"github.com/BESTSELLER/{{cookiecutter.name}}/logger"
)

func main() {
	config.LoadEnvConfig()
	logger.Init()

	api.SetupRouter()
}
