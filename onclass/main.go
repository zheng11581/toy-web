package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

func signUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReaJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	server := NewHttpServer("test-server")
	//server.Route("/", home)
	//server.Route("/user", user)
	//server.Route("/user/create", createUser)
	server.Route("POST", "/user/signup", signUp)
	//server.Route("/order", order)
	server.Start(":8080")
}
