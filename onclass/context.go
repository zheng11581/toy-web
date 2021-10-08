package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (ctx *Context) ReaJson(req interface{}) error {
	r := ctx.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	return nil
}

func (ctx *Context) WriteJson(code int, resp interface{}) error {
	ctx.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = ctx.W.Write(respJson)
	return err
}

func (ctx *Context) WriteOkJson(resp interface{}) error {
	return ctx.WriteJson(http.StatusOK, resp)
}

func (ctx *Context) WriteSystemErrorJson(resp interface{}) error {
	return ctx.WriteJson(http.StatusInternalServerError, resp)
}

func (ctx *Context) BadRequestJson(resp interface{}) error {
	return ctx.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer,
		R: request,
	}
}
