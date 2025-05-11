package system

import (
	"go-server/config"
	"go-server/internal/console"
	"go-server/utils/errors/errorsApp"
	"go-server/utils/os"
)

func CheckSystemRequirements() error {
	done := console.PrintStatusBar("Check system requirements")

	if !os.IsSupportedOs(&config.SupportOs) {
		done <- false
		return errorsApp.ErrOsNotSupported
	}

	if os.GetOS() == os.MacOs || os.GetOS() == os.Linux {
		if !os.IsRoot() && !os.CheckWifiCmd() {
			done <- false
			return errorsApp.ErrNeedSudo
		}
	}

	done <- true
	return nil
}
