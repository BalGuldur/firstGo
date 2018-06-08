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
	Name string
	id   string
}

func (device *Device) Save() {
	if device.id == "" {
		device.id = randomId()
	}
	store[device.id] = *device
}

func (device Device) Delete() error {
	if _, ok := store[device.id]; ok {
		delete(store, device.id)
		return nil
	} else {
		return errors.New("not found")
	}
}

// func Find(id string) (Device, error) {
// 	if dev, ok := store[id]; ok {
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
// 		id:   m["id"].(string),
// 	}
// 	return dev
// }
