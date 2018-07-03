package app

import (
	"devices/app/websock"
	"log"
	"net/http"
)

func Start() {
	websock.Start()
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
