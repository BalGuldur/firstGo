package websock

import (
	// "bytes"
	"../processor"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var (
// newline = []byte{'\n'}
// space   = []byte{' '}
)

func createReadPipe(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.SetReadLimit(maxMessageSize)
	// conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		// _, message, err := conn.ReadMessage()
		var v = processor.Request{}
		err := conn.ReadJSON(&v)
		fmt.Println("new message")
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			} else {
				log.Printf("error")
			}
		}
		// message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println(v)
		// processor.Exec(v)
	}
}
