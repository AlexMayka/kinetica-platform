package errorsWifi

import (
	"errors"
	errors2 "go-server/utils/errors"
)

var (
	ErrWiFiConnect    = errors.New("failed to connect to WiFi")
	ErrDeviceNotFound = errors.New("device not found")
	ErrWifiTuring     = errors.New("turing failed")
)

func init() {
	errors2.Registration(ErrWiFiConnect, errors2.ErrorResponse{Code: 3001, Message: "Could not connect to Wi-Fi", Tag: "wifi_connect"})
	errors2.Registration(ErrDeviceNotFound, errors2.ErrorResponse{Code: 3002, Message: "Could not find Wi-Fi", Tag: "wifi_search"})
	errors2.Registration(ErrWifiTuring, errors2.ErrorResponse{Code: 3003, Message: "Unable to turn on Wifi", Tag: "wifi_turing"})
}
