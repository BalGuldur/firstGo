package app

import (
	"math/rand"
	"strconv"
)

type idType string

var deviceStore = make(map[idType]Device)

func randomId() idType {
	return idType(strconv.Itoa(rand.Int()))
}

type Device struct {
	name string
	id   idType
}

func newDevice(name string) Device {
	return Device{name: name, id: randomId()}
}

func allDevices() map[idType]Device {
	return deviceStore
}

func (dev Device) save() error {
	deviceStore[idType(dev.id)] = dev
	return nil
}

func (dev Device) delete() error {
	delete(deviceStore, dev.id)
	return nil
}

func destroyAll() error {
	deviceStore = make(map[idType]Device)
	return nil
}

// AddDevice action of Device controller =)
func AddDevice(params actionParams) interface{} {
	dev := newDevice(params["name"])
	dev.save()
	return dev
}

func AllDevices(params actionParams) interface{} {
	return allDevices()
}

func RemoveAllDevices(params actionParams) interface{} {
	return destroyAll()
}

func RemoveDevice(params actionParams) interface{} {
	return deviceStore[idType(params["id"])].delete()
}
