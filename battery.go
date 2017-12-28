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
// battery:
//-----------------------------------------------------------------------------

func battery(i *item) {

	var eNow, eFull, perc int
	const path = "/sys/class/power_supply/BAT0/"

	for {

		// Energy now:
		energyNow, err := ioutil.ReadFile(path + "energy_now")
		if err != nil {
			eNow = -1
		} else if _, err := fmt.Sscanf(string(energyNow), "%d", &eNow); err != nil {
			eNow = -1
		}

		// Energy full:
		energyFull, err := ioutil.ReadFile(path + "energy_full")
		if err != nil {
			eFull = -1
		} else if _, err := fmt.Sscanf(string(energyFull), "%d", &eFull); err != nil {
			eFull = -1
		}

		// Percentage:
		if (eNow == -1) || (eFull == -1) {
			perc = -1
		} else {
			perc = eNow * 100 / eFull
		}

		// Set and sleep:
		i.data = softPurple + iconBattery + strconv.Itoa(perc) + "%"
		time.Sleep(time.Minute)
	}
}
