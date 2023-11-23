package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = &websocket.Upgrader{
	//定义读写缓冲区大小
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	//校验请求
	CheckOrigin: func(r *http.Request) bool {
		//如果不是get请求，返回错误
		if r.Method != "GET" {
			fmt.Println("请求方式错误")
			return false
		}
		//如果路径中不包括chat，返回错误
		if r.URL.Path != "/chat" {
			fmt.Println("请求路径错误")
			return false
		}
		//还可以根据其他需求定制校验规则
		return true
	},
}

func WsHandle(writer http.ResponseWriter, req *http.Request) {
	conn, err := upGrader.Upgrade(writer, req, nil)
	if err != nil {
		fmt.Println("ws 连接升级失败")
	}
	fmt.Println(conn)
}
