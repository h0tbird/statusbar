package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import "time"

//-----------------------------------------------------------------------------
// updates:
//-----------------------------------------------------------------------------

func updates(i *item) {
	for {
		i.data = softWhite + iconUpdates + "0"
		time.Sleep(5 * time.Minute)
	}
}