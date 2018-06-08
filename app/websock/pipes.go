package websock

import (
	"../processor"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func createReadPipe(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.SetReadLimit(maxMessageSize)
	// conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := conn.ReadMessage()
		// var request = processor.Request{}
		// err := conn.ReadJSON(&request)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			} else {
				log.Printf("error")
			}
		}
		processor.Proceed(message)
	}
}
