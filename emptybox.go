package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xcursor"
	"github.com/BurntSushi/xgbutil/xevent"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// connect to X
	X, err := xgbutil.NewConn()
	checkError(err)
	defer X.Conn().Close()

	cursor := createCursor(X, xcursor.LeftPtr)

	// dump X connection obj
	XC := X.Conn()

	// assign default cursor to "the default invisible root window"
	cookie := xproto.ChangeWindowAttributesChecked(XC, X.RootWin(),
		xproto.CwBackPixmap|xproto.CwEventMask|xproto.CwCursor,
		[]uint32{
			xproto.BackPixmapParentRelative,
			xproto.EventMaskButtonPress |
				xproto.EventMaskButtonRelease |
				xproto.EventMaskButtonMotion |
				xproto.EventMaskPointerMotion,
			uint32(cursor),
		})
	err = cookie.Check()
	checkError(err)

	//	win, err := xwindow.Generate(X)
	//	checkError(err)
	//	win.Create(X.RootWin(), 0, 0, 500, 500,
	//		xproto.CwBackPixel|xproto.CwCursor,
	//		0xffffffff, uint32(cursor))
	//	win.Map()

	// setting up event handling
	keybind.Initialize(X)

	keybind.KeyPressFun(
		func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
			drawMenu(X, 0, 10, 12.0)
			fmt.Println("-----------------")
			fmt.Println("WORKS")
		}).Connect(X, X.RootWin(), "Mod4-Control-Shift-t", true)
}
