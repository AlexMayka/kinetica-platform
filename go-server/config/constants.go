package config

import "go-server/utils/os"

const (
	AppVersion = "v0.1.0"
	PathConfig = "settings/settings.json"
)

var SupportOs = []os.OS{os.MacOs}
