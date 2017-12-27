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
	softWhite  = "\x02"
	softOrange = "\x03"
	leftArrow  = "\xEE\x86\xAC"
	wallClock  = "\xEE\x80\x95"
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

func printFoo(i *item) {
	for _ = range time.NewTicker(2 * time.Second).C {
		i.data = "y"
	}
}

func printBar(i *item) {
	for _ = range time.NewTicker(4 * time.Second).C {
		i.data = "z"
	}
}

func dateTime(i *item) {
	for _ = range time.NewTicker(1 * time.Second).C {
		i.data = time.Now().Format(
			softWhite + wallClock + "Mon _2 Jan " + leftArrow + "15:04:05")
	}
}

//-----------------------------------------------------------------------------
// Main loop:
//-----------------------------------------------------------------------------

func main() {

	// Initialize the structure:
	items := []*item{
		&item{true, "", printFoo},
		&item{true, "", printBar},
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
		setStatus(strings.Join(status, " "+softOrange+"| ") + "       ")
	}
}
