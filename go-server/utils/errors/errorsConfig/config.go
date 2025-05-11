package errorsConfig

import (
	"errors"
	errors2 "go-server/utils/errors"
)

var (
	ErrConfigOpen       = errors.New("cannot open config file")
	ErrConfigRead       = errors.New("cannot read config file")
	ErrConfigParse      = errors.New("invalid config format")
	ErrMissingBLE       = errors.New("missing BLE CommandCharUUID")
	ErrMissingTCPServer = errors.New("missing TCP server")
	ErrMissingTCPPort   = errors.New("missing TCP port")
	ErrMissingTCP       = errors.New("missing TCP block")
)

func init() {
	errors2.Registration(ErrConfigOpen, errors2.ErrorResponse{Code: 1001, Message: "Cannot open configuration file", Tag: "config_open"})
	errors2.Registration(ErrConfigRead, errors2.ErrorResponse{Code: 1002, Message: "Failed to read configuration file", Tag: "config_rea"})
	errors2.Registration(ErrConfigParse, errors2.ErrorResponse{Code: 1003, Message: "Configuration file format is invalid", Tag: "config_parse"})
	errors2.Registration(ErrMissingBLE, errors2.ErrorResponse{Code: 1004, Message: "Missing BLE UUID", Tag: "config_ble"})
	errors2.Registration(ErrMissingTCPServer, errors2.ErrorResponse{Code: 1005, Message: "Missing TCP server", Tag: "config_tcp_server"})
	errors2.Registration(ErrMissingTCPPort, errors2.ErrorResponse{Code: 1006, Message: "Missing TCP port", Tag: "config_tcp_port"})
	errors2.Registration(ErrMissingTCP, errors2.ErrorResponse{Code: 1007, Message: "Missing TCP block", Tag: "config_tcp_block"})
}
