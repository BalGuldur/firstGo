package websock

import (
	"devices/app/processor"
	"github.com/gorilla/websocket"
)

type channel struct {
	connections []websocket.Conn
}

func (channel *channel) addConn(conn *websocket.Conn) {
	channel.connections = append(channel.connections, *conn)
}

func (channel *channel) send(response processor.Response, selfConn *websocket.Conn) {
	for _, conn := range channel.connections {
		write(&conn, response)
	}
}
