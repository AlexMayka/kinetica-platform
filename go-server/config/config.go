package config

import (
	"encoding/json"
	"go-server/internal/console"
	"go-server/utils/errors/errorsConfig"
	"io"
	"os"
)

type BLE struct {
	CommandCharUUID *string `json:"command_char_uuid"`
}

type TCP struct {
	Server *string `json:"server"`
	Port   *string `json:"port"`
}

type Config struct {
	BLE *BLE `json:"BLE"`
	TCP *TCP `json:"TCP"`
}

var config = &Config{}

func readData(path string) (*[]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errorsConfig.ErrConfigOpen
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errorsConfig.ErrConfigRead
	}
	return &data, err
}

func (c *Config) Validate() error {
	if c.BLE == nil {
		return errorsConfig.ErrMissingBLE
	}
	if c.BLE.CommandCharUUID == nil || len(*c.BLE.CommandCharUUID) == 0 {
		return errorsConfig.ErrMissingBLE
	}

	if c.TCP == nil {
		return errorsConfig.ErrMissingTCP
	}
	if c.TCP.Server == nil || len(*c.TCP.Server) == 0 {
		return errorsConfig.ErrMissingTCPServer
	}
	if c.TCP.Port == nil || len(*c.TCP.Port) == 0 {
		return errorsConfig.ErrMissingTCPPort
	}

	return nil
}

func Load(path string) error {
	done := console.PrintStatusBar("Initializing the config")
	data, err := readData(path)
	if err != nil {
		done <- false
		return err
	}

	if err := json.Unmarshal(*data, config); err != nil {
		done <- false
		return err
	}

	if err := config.Validate(); err != nil {
		done <- false
		return err
	}

	done <- true
	return nil
}

func GetConfig() *Config {
	return config
}

func GetBLE() *BLE {
	return config.BLE
}

func GetTCP() *TCP {
	return config.TCP
}
