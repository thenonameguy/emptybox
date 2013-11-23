package main

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xcursor"
)

// Creates a cursor, and returns it.
// Type: see consts in xcursor
func createCursor(X *xgbutil.XUtil, Type uint16) xproto.Cursor {
	cursor, err := xcursor.CreateCursor(X, Type)
	checkError(err)
	return cursor
}
