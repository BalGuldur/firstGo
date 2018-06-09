package app

import (
	"./websock"
	"log"
	"net/http"
)

func Start() {
	websock.Start()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
