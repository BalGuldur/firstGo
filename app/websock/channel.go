package websock

import "github.com/gorilla/websocket"

type channel struct {
	connections []websocket.Conn
}

func (channel *channel) addConn(conn *websocket.Conn) {
	channel.connections = append(channel.connections, *conn)
}
