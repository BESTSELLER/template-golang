package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/{{cookiecutter.destination.git.owner + "/" + cookiecutter.destination.git.name}}/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

// prometheusMiddleware implements mux.MiddlewareFunc.
func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

// SetupRouter initializes the API routes
func SetupRouter() {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	loggedRouter := router.PathPrefix("/").Subrouter()

	loggedRouter.Use(prometheusMiddleware)
	loggedRouter.Use(func(next http.Handler) http.Handler { return handlers.CombinedLoggingHandler(os.Stdout, next) })

	loggedRouter.HandleFunc("/", defaultHandler).Methods("GET")
	loggedRouter.HandleFunc("/endpoint1", endpoint1).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.EnvVars.Port), router))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the default handler")
}
func endpoint1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from endpoint1")
}
