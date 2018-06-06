package app

type actionParams map[string]string
type actionFunc func(actionParams) interface{}

var actions = make(map[string]actionFunc)

func init() {
	actions["device#all"] = AllDevices
	actions["device#add"] = AddDevice
	actions["device#remove_all"] = RemoveAllDevices
	actions["device#remove"] = RemoveDevice
}
