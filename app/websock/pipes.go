package websock

import (
	"../processor"
	"github.com/gin-gonic/gin/json"
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
		var response = processor.Proceed(message)
		if response.Public {
			broadcast.send(response, conn)
		} else {
			write(conn, response)
		}
	}
}

func write(conn *websocket.Conn, resp processor.Response) {
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	byteResp, _ := json.Marshal(resp)
	w.Write(byteResp)
	if err := w.Close(); err != nil {
		return
	}
}
