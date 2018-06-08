package websock

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var broadcast = channel{connections: make([]websocket.Conn, 0)}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// TODO: write with pipe for safe
	broadcast.addConn(conn)

	// go createWritePipe(conn)
	go createReadPipe(conn)
}

func Start() {
	http.HandleFunc("/ws", serveWs)
}
