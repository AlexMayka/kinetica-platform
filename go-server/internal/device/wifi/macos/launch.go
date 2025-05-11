package macos

import (
	"go-server/utils/errors/errorsWifi"
	"os/exec"
	"regexp"
	"strings"
)

type WifiMacOs struct{}

func getDevice() (*string, error) {
	cmd := exec.Command("networksetup", "-listallhardwareports")
	out, err := cmd.Output()

	if err != nil {
		return nil, errorsWifi.ErrDeviceNotFound
	}

	pattern := `Wi-Fi\s*Device:\s*(?<DvName>[\S]+)\s`
	reg, _ := regexp.Compile(pattern)
	result := reg.FindSubmatch(out)

	if len(result) != 2 {
		return nil, errorsWifi.ErrDeviceNotFound
	}

	device := string(result[1])

	return &device, nil
}

func isActive(device *string) bool {
	cmd := exec.Command("networksetup", "-getairportpower", *device)
	out, err := cmd.Output()

	if err != nil {
		return false
	}

	return strings.Contains(string(out), "On")
}

func turning(device *string) error {
	if isActive(device) {
		return nil
	}
	cmd := exec.Command("networksetup", "-setairportpower", *device, "on")
	_, err := cmd.Output()

	if err != nil {
		return errorsWifi.ErrWifiTuring
	}

	return nil
}

func (wifi *WifiMacOs) Launch() error {
	device, err := getDevice()
	if err != nil {
		return err
	}

	if err := turning(device); err != nil {
		return err
	}

	return nil
}
