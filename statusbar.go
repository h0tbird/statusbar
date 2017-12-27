package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import (
	"strings"
	"time"
)

//-----------------------------------------------------------------------------
// Constants:
//-----------------------------------------------------------------------------

const (

	// Colors:
	softWhite  = "\x02"
	softOrange = "\x03"
	softPurple = "\x06"

	// Icons:
	iconLeftArrow = "\xEE\x86\xAC"
	iconWallClock = "\xEE\x80\x95"
	iconBattery   = "\xEE\x80\xB3"
	iconUpdates   = "\xEE\x80\x8E"
)

//-----------------------------------------------------------------------------
// Item structure:
//-----------------------------------------------------------------------------

type item struct {
	show bool
	data string
	fn   func(*item)
}

func (i *item) runFunc() {
	if i.fn != nil {
		i.fn(i)
	}
}

//-----------------------------------------------------------------------------
// Item functions:
//-----------------------------------------------------------------------------

func updates(i *item) {
	for _ = range time.NewTicker(5 * time.Minute).C {
		i.data = softWhite + iconUpdates + "0"
	}
}

func battery(i *item) {
	for _ = range time.NewTicker(time.Minute).C {
		i.data = softPurple + iconBattery + "100%"
	}
}

func dateTime(i *item) {
	for _ = range time.NewTicker(1 * time.Second).C {
		i.data = time.Now().Format(softWhite + iconWallClock +
			"Mon _2 Jan " + iconLeftArrow + "15:04:05")
	}
}

//-----------------------------------------------------------------------------
// Main loop:
//-----------------------------------------------------------------------------

func main() {

	// Initialize the structure:
	items := []*item{
		&item{true, "", updates},
		&item{true, "", battery},
		&item{true, "", dateTime},
	}

	// Start each item's logic:
	for _, v := range items {
		go v.runFunc()
	}

	// Refresh the status bar:
	for _ = range time.NewTicker(time.Second).C {
		status := []string{}
		for _, item := range items {
			if item.show {
				status = append(status, item.data)
			}
		}
		setStatus(strings.Join(status, " "+softOrange+"| ") + "          ")
	}
}
