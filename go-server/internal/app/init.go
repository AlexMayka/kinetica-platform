package app

import (
	"go-server/config"
	"go-server/internal/device/ble"
	"go-server/internal/device/wifi"
	"go-server/internal/system"
	"go-server/utils/errors"
)

func AppInit() error {
	if err := system.CheckSystemRequirements(); err != nil {
		errors.PrintError(err)
		return err
	}

	if err := config.Load(config.PathConfig); err != nil {
		errors.PrintError(err)
		return err
	}

	if err := ble.InitBLE(); err != nil {
		errors.PrintError(err)
		return err
	}

	if err := wifi.InitWifi(); err != nil {
		errors.PrintError(err)
		return err
	}

	return nil
}
