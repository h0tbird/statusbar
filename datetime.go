package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import "time"

//-----------------------------------------------------------------------------
// dateTime:
//-----------------------------------------------------------------------------

func dateTime(i *item) {
	for {
		i.data = time.Now().Format(softWhite + iconWallClock +
			"Mon _2 Jan " + iconLeftArrow + "15:04:05")
		time.Sleep(time.Second)
	}
}