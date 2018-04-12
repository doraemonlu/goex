package ipc

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/testdata/my_test"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer(server)
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0) //0说明没有缓冲区

	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req) //反序列化，从json到struct
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)
			b, err = json.Marshal(resp) //序列化，从struct到json

			c <- string(b)
		}

		fmt.Println("Session closed.")
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session
}