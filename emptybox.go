package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xcursor"
	"os"
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
	/*pingBefore, pingAfter, pingQuit := xevent.MainPing(X)
	EVENTLOOP:
		for {
			select {
			case <-pingBefore:
				// Wait for event processing to finish.
				<-pingAfter
			case val := <-someOtherChannel:
				// do some work with val
			case <-pingQuit:
				break EVENTLOOP
			}
		}*/
	drawMenu(X, 30, 30, 12.0)
	for {
		ev, xerr := XC.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			os.Exit(1)
		}

		if ev != nil {
			fmt.Printf("Event: %s\n", ev)
		}
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}
