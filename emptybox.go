package main

import (
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xcursor"
	//"github.com/BurntSushi/xgbutil/xevent"
)

func main() {
	// connect to X
	X, err := xgbutil.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer X.Conn().Close()

	// create cursor
	_, err = xcursor.CreateCursor(X, xcursor.Circle)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create an xgb_conn for later use 
	xgb_conn := X.Conn()

	// create a sample window
	wid, _ := xproto.NewWindowId(xgb_conn)
	screen := xproto.Setup(xgb_conn).DefaultScreen(xgb_conn)
	xproto.CreateWindow(xgb_conn, screen.RootDepth, wid, screen.Root,
		0, 0, 500, 500, 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel | xproto.CwEventMask,
		[]uint32{ // values must be in the order defined by the protocol
			0xffffffff,
			xproto.EventMaskStructureNotify |
			xproto.EventMaskKeyPress |
			xproto.EventMaskKeyRelease})

	xproto.MapWindow(xgb_conn, wid)

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
		ev, xerr := xgb_conn.WaitForEvent()
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
