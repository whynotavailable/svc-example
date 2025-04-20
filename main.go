package main

// An example application with a health check and a simple RPC function.
// The RPC container requires an authorization header with a value of `good`.

import (
	"net/http"

	"example/middleware"
	"example/rpcs"

	"github.com/whynotavailable/svc"
)

func httpRoutes() {
	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
}

func setupRpcs() {
	container := svc.NewRpcContainer()

	container.AddFunction(rpcs.GetWeatherKey, rpcs.GetWeather).
		BodyType(rpcs.GetWeatherRequest{}).
		Meta("hi", "fred")

	container.GenerateDocs()

	var handler http.Handler = &middleware.AuthMiddleware{
		Inner: container,
	}

	svc.SetupContainer(http.DefaultServeMux, "/rpc", handler)
}

func main() {
	httpRoutes()
	setupRpcs()

	handler := svc.NewLoggingMiddleware(http.DefaultServeMux)

	http.ListenAndServe("0.0.0.0:4321", handler)
}
