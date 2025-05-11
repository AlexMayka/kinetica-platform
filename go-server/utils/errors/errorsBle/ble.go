package errorsBle

import (
	"errors"
	errors2 "go-server/utils/errors"
)

var (
	ErrBLEDisabled = errors.New("failed to initialize BLE")
)

func init() {
	errors2.Registration(ErrBLEDisabled, errors2.ErrorResponse{Code: 2001, Message: "BLE is disabled. Please turn it on and try again", Tag: "ble_init"})
}
