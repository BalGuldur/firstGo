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
	Action  string `json:"action"`
	Success bool   `json:"success"`
	Public  bool   `json:"-"`
	// Device  *devices.Device  `json:"device,omitempty"`
	Devices []devices.Device `json:"devices,omitempty"`
}

func Proceed(raw []byte) Response {
	var request = Request{}
	if err := json.Unmarshal(raw, &request); err != nil {
		// TODO: Return error and send in response message
		return Response{}
	}
	var response = Response{
		Success: true,
		Public:  false,
		Action:  request.Action,
	}
	switch request.Action {
	case "devices.add":
		var dev, suc, _ = request.Device.Save()
		response.Devices = []devices.Device{dev}
		response.Success = suc
		response.Public = true
	case "devices.all":
		response.Devices = devices.All()
	case "devices.delete":
		var dev, suc, _ = request.Device.Delete()
		response.Devices = []devices.Device{dev}
		response.Success = suc
		response.Public = true
	case "devices.get":
		var dev, _ = devices.Find(request.Device.Id)
		response.Devices = []devices.Device{dev}
	}
	return response
}
