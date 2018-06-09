package processor

import (
	"../devices"
	"encoding/json"
)

type Request struct {
	Action string         `json:"action"`
	Device devices.Device `json:"device"`
}

type Response struct {
	Action  string           `json:"action"`
	Success bool             `json:"success"`
	Public  bool             `json:"-"`
	Device  *devices.Device  `json:"device,omitempty"`
	Devices []devices.Device `json:"devices,omitempty"`
}

func Proceed(raw []byte) Response {
	var request = Request{}
	if err := json.Unmarshal(raw, &request); err != nil {
		// TODO: Return error and send in response message
		panic(err)
	}
	var response = Response{
		Success: true,
		Public:  false,
		Action:  request.Action,
	}
	switch request.Action {
	case "device.add":
		var dev, suc, _ = request.Device.Save()
		response.Device = &dev
		response.Success = suc
		response.Public = true
	case "device.all":
		response.Devices = devices.All()
	case "device.delete":
		var dev, suc, _ = request.Device.Delete()
		response.Device = &dev
		response.Success = suc
		response.Public = true
	case "device.get":
		var dev, _ = devices.Find(request.Device.Id)
		response.Device = &dev
	}
	return response
}
