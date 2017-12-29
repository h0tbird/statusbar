package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

//-----------------------------------------------------------------------------
// Constants:
//-----------------------------------------------------------------------------

const (

	// Icons:
	iconBattery90 = "\xEE\x80\xB3"
	iconBattery63 = "\xEE\x80\xB6"
	iconBattery37 = "\xEE\x80\xB5"
	iconBattery10 = "\xEE\x80\xB4"
	iconPlugged   = "\xEE\x81\x81"
)

//-----------------------------------------------------------------------------
// battery:
//-----------------------------------------------------------------------------

func battery(i *item) {

	var eNow, eFull, plug, perc int
	const path = "/sys/class/power_supply"

	for {

		// Energy now:
		energyNow, err := ioutil.ReadFile(path + "/BAT0/energy_now")
		if err != nil {
			eNow = -1
		} else if _, err := fmt.Sscanf(string(energyNow), "%d", &eNow); err != nil {
			eNow = -1
		}

		// Energy full:
		energyFull, err := ioutil.ReadFile(path + "/BAT0/energy_full")
		if err != nil {
			eFull = -1
		} else if _, err := fmt.Sscanf(string(energyFull), "%d", &eFull); err != nil {
			eFull = -1
		}

		// Plugged:
		plugged, err := ioutil.ReadFile(path + "/AC/online")
		if err != nil {
			perc = -1
		} else if _, err := fmt.Sscanf(string(plugged), "%d", &plug); err != nil {
			perc = -1
		}

		// Percentage:
		if (eNow == -1) || (eFull == -1) {
			perc = -1
		} else {
			perc = eNow * 100 / eFull
		}

		// Icon:
		icon := iconPlugged
		if plug == 0 {
			switch {
			case perc >= 90:
				icon = iconBattery90
			case perc >= 63:
				icon = iconBattery63
			case perc >= 37:
				icon = iconBattery37
			case perc >= 10:
				icon = iconBattery10
			}
		}

		// Set and sleep:
		i.data = softPurple + icon + strconv.Itoa(perc) + "%"
		time.Sleep(time.Minute)
	}
}
