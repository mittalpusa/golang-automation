package appsaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/sclevine/agouti/appium"
)

// Device global variable
var Device *appium.Device

// StartDriver : start android driver
func (s *Page) StartDriver() error {
	return s.driver().Start()
}

// NewDevice : create new android device
func (s *Page) NewDevice() *appium.Device {
	var err error

	Device, err = s.driver().NewDevice()
	helper.LogPanicln(err)

	return Device
}
