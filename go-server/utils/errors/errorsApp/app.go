package errorsApp

import (
	"errors"
	errors2 "go-server/utils/errors"
)

var (
	ErrOsNotSupported      = errors.New("OS not supported")
	ErrNeedSudo            = errors.New("not root")
	ErrMissingCommandError = errors.New("missing command")
)

func init() {
	errors2.Registration(ErrOsNotSupported, errors2.ErrorResponse{Code: 0001, Message: "OS not supported", Tag: "os_not_supported"})
	errors2.Registration(ErrNeedSudo, errors2.ErrorResponse{Code: 0002, Message: "Must run as root (sudo)", Tag: "no_sudo"})
	errors2.Registration(ErrMissingCommandError, errors2.ErrorResponse{Code: 0003, Message: "Can't work with cmd", Tag: "no_sudo"})
}
