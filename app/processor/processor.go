package processor

import (
	"../devices"
	"encoding/json"
	"fmt"
)

type Request struct {
	Action string         `json:"action"`
	Device devices.Device `json:"device"`
}

// func (req Request) Run() {
// 	if strings.Contains(req.Action, "device.") {
// 		var model = req.Model.(devices.Device)
// 	}
// 	switch req.Action {
// 	case "device.add":
//
// 	}
// }

// func Exec(request map[string]interface{}) {
// }
func Proceed(raw []byte) {
	var request = Request{}
	if err := json.Unmarshal(raw, &request); err != nil {
		// TODO: Return error and send in response message
		panic(err)
	}
	switch request.Action {
	case "device.add":
		request.Device.Save()
	case "device.all":
		fmt.Println(devices.All())
	case "device.delete":
		request.Device.Delete()
	}
}
