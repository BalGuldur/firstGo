package app

import "fmt"
import "./devices"
import (
	"./websock"
	"log"
	"net/http"
)

func Start() {
	dev1 := devices.Device{Name: "test dev"}
	fmt.Println(dev1)
	dev1.Save()
	fmt.Println(devices.All())
	dev1.Name = "other name"
	fmt.Println(dev1)
	fmt.Println(devices.All())
	dev1.Delete()
	fmt.Println(devices.All())
	websock.Start()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
