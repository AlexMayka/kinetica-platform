package wifi

import (
	"errors"
	"fmt"
	"go-server/internal/console"
	"go-server/internal/device/wifi/macos"
	"go-server/utils/os"
)

type Wifi interface {
	Launch() error
}

var wifi Wifi

func InitWifi() error {
	done := console.PrintStatusBar("Initializing the Wifi")
	switch os.GetOS() {
	case os.MacOs:
		wifi = &macos.WifiMacOs{}
		if err := wifi.Launch(); err != nil {
			done <- false
			return err
		}
		done <- true
		return nil

	default:
		done <- false
		return errors.New(fmt.Sprintf("Unsupported OS: %s", os.GetOS()))
	}
}

func GetWifi() Wifi {
	return wifi
}
