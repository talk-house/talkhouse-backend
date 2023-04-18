package socket

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type Message struct {
	Msg string `json:"msg"`
}

func StartWebsocket(ws *websocket.Conn) {
	fmt.Println("Socket is Established")

	for {
		var msg Message

		websocket.JSON.Receive(ws, &msg)
		fmt.Println("Recieved Message:", msg.Msg)

		response := Message{Msg: "Thanks for sending Us a messsage"}
		websocket.JSON.Send(ws, response)
	}
}
