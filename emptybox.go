package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xcursor"
	"runtime"
	//"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// connect to X
	X, err := xgbutil.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer X.Conn().Close()

	// create a cursor
	cursor, err := xcursor.CreateCursor(X, xcursor.LeftPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

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
	if err != nil {
		fmt.Println(err)
		return
	}

	win, err := xwindow.Generate(X)
	if err != nil {
		fmt.Println(err)
		return
	}
	win.Create(X.RootWin(), 0, 0, 500, 500,
		xproto.CwBackPixel|xproto.CwCursor,
		0xffffffff, uint32(cursor))
	win.Map()

	// create a sample window
	//wid, _ := xproto.NewWindowId(XC)
	//screen := xproto.Setup(XC).DefaultScreen(XC)
	//xproto.CreateWindow(XC, screen.RootDepth, wid, screen.Root,
	//	0, 0, 500, 500, 0,
	//	xproto.WindowClassInputOutput, screen.RootVisual,
	//	xproto.CwBackPixel | xproto.CwEventMask,
	//	[]uint32{ // values must be in the order defined by the protocol
	//		0xffffffff,
	//		xproto.EventMaskStructureNotify |
	//		xproto.EventMaskKeyPress |
	//		xproto.EventMaskKeyRelease})

	//xproto.MapWindow(XC, wid)

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

	for {
		ev, xerr := XC.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			fmt.Printf("Event: %s\n", ev)
		}
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}
