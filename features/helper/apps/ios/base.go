package ios

import (
	"fmt"
	"log"

	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

/*Driver global variable*/
var Driver *appium.WebDriver

/*Device global variable*/
var Device *appium.Device

/*DriverConnect for ios*/
func DriverConnect() error {
	p := properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-ios.properties", properties.UTF8)

	options := appium.Desired(agouti.Capabilities{
		"automationName":    p.MustGetString("automationName"),
		"platformName":      p.MustGetString("platformName"),
		"newCommandTimeout": p.MustGetString("newCommandTimeout"),
		"app":               p.MustGetString("app"),
		"platformVersion":   p.MustGetString("platformVersion"),
		"deviceName":        p.MustGetString("deviceName"),
		"fullReset":         p.MustGetString("fullReset"),
		"autoAcceptAlerts":  p.MustGetString("autoAcceptAlerts"),
		"xcodeOrgId":        p.MustGetString("xcodeOrgId"),
		"xcodeSigningId":    p.MustGetString("xcodeSigningId"),
		"udid":              p.MustGetString("udid"),
		"xcodeConfigfile":   p.MustGetString("xcodeConfigfile"),
		"useNewWDA":         p.MustGetString("useNewWDA"),
		"showXcodeLog":      p.MustGetString("showXcodeLog"),
	})

	Driver = appium.New(options)

	return nil
}

/*OpenApps start and create new device ios*/
func OpenApps() error {
	var err error

	if err := Driver.Start(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	if Device, err = Driver.NewDevice(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
