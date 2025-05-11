package os

import (
	"os"
	"os/exec"
	"runtime"
	"slices"
)

type OS string

const (
	Windows OS = "windows"
	Linux   OS = "linux"
	MacOs   OS = "darwin"
)

func IsSupportedOs(supportOs *[]OS) bool {
	return slices.Contains(*supportOs, GetOS())
}

func GetOS() OS {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "linux":
		return Linux
	default:
		return MacOs
	}
}

func IsRoot() bool {
	return os.Getuid() == 0
}

func CheckWifiCmd() bool {
	switch GetOS() {
	case MacOs:
		return CheckCMD("networksetup")
	case Windows:
		return true
	default:
		return false
	}
}

func CheckCMD(cmd string) bool {
	if _, err := exec.LookPath(cmd); err != nil {
		return false
	}
	return true
}
