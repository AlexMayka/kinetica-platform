package ble

import (
	"go-server/config"
	"tinygo.org/x/bluetooth"
)

var (
	ble *BLE
	cnf *config.BLE
)

type BLE struct {
	charUUID string
	adapter  *bluetooth.Adapter
	dev      *bluetooth.ScanResult
	client   *bluetooth.Device
}
