package ipc


import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request(methd, params)

	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	client.conn <- string(b) // 向通道发送请求
	str := <-client.conn // 接收通道响应

	var resp1 Response
	json.Unmarshal([]byte(str), &resp1)

	resp = &resp1

	return
}

func (client *IpcClient) Close() {
	client.conn <- "ClOSE"
}