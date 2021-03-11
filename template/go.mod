module github.com/{{cookiecutter.destination.git.owner + "/" + cookiecutter.destination.git.name}}

go 1.15

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/client_golang v1.9.0
	github.com/rs/zerolog v1.20.0
)
