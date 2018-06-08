package devices

import (
	"errors"
	"math/rand"
	"strconv"
)

var store = make(map[string]Device)

func randomId() string {
	return string(strconv.Itoa(rand.Int()))
}

type Device struct {
	Name string `json:"name"`
	Id   string `json:"Id"`
}

func (device *Device) Save() {
	if device.Id == "" {
		device.Id = randomId()
	}
	store[device.Id] = *device
}

func (device Device) Delete() error {
	if _, ok := store[device.Id]; ok {
		delete(store, device.Id)
		return nil
	} else {
		return errors.New("not found")
	}
}

// func Find(Id string) (Device, error) {
// 	if dev, ok := store[Id]; ok {
// 		return dev, nil
// 	} else {
// 		return dev, errors.New("not found")
// 	}
// }

func All() (result []Device) {
	for _, dev := range store {
		result = append(result, dev)
	}
	return
}

// func New() (dev Device) {
// 	return
// }

// func FromMap(m map[string]interface{}) Device {
// 	dev := Device{
// 		Name: m["name"].(string),
// 		Id:   m["Id"].(string),
// 	}
// 	return dev
// }
