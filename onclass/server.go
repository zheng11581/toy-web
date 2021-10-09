package main

import (
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedOnMap
}

func (s *sdkHttpServer) Route(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc
}

func (s sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return err
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
