package main

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.Handle(pattern, handlerFunc)
}

func (s sdkHttpServer) Start(address string) error {
	err := http.ListenAndServe(address, nil)
	return err
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
