package main

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}
