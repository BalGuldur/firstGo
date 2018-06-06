package app

import "fmt"

var (
	emptyParams = make(actionParams)
)

func Start() {
	params := make(actionParams)
	params["name"] = "test dev"
	fmt.Println(actions["device#all"](emptyParams))
	actions["device#add"](params)
	dev := actions["device#add"](params).(Device)
	fmt.Println(dev)
	fmt.Println(actions["device#all"](emptyParams))
	params = make(actionParams)
	params["id"] = string(dev.id)
	fmt.Println(actions["device#remove"](params))
	fmt.Println(actions["device#all"](emptyParams))
}
