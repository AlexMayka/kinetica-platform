package main

import (
	"go-server/config"
	"go-server/internal/app"
	"go-server/internal/console/banner"
	"go-server/utils/errors"
)

func main() {
	banner.LaunchAPP(config.AppVersion)
	err := app.AppInit()
	if err != nil {
		errors.PrintError(err)
		return
	}

}
