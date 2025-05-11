package ble

import (
	"go-server/config"
	"go-server/internal/console"
	errBle "go-server/utils/errors/errorsBle"
	"tinygo.org/x/bluetooth"
)

func InitBLE() error {
	if ble != nil {
		return nil
	}

	done := console.PrintStatusBar("Initializing the BLE")
	adapter := bluetooth.DefaultAdapter
	if err := adapter.Enable(); err != nil {
		done <- false
		return errBle.ErrBLEDisabled
	}

	cnf = config.GetBLE()

	ble = &BLE{
		adapter:  adapter,
		charUUID: *cnf.CommandCharUUID,
	}
	done <- true
	return nil
}
