package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"image"
	"os"
)

func drawMenu(X *xgbutil.XUtil, pos_x, pos_y int, size float64) {
	// background color of the canvas
	var bg = xgraphics.BGRA{B: 0xff, G: 0x66, R: 0x33, A: 0xff}

	// color of the text
	var fg = xgraphics.BGRA{B: 0xff, G: 0xff, R: 0xff, A: 0xff}

	// select font
	fontReader, err := os.Open("/usr/share/fonts/dejavu/DejaVuSansMono.ttf")
	checkError(err)
	defer fontReader.Close()

	// parse font
	font, err := xgraphics.ParseFont(fontReader)
	checkError(err)

	msg := "it works"
	// get proper width and height of the 1 line text
	_, height := xgraphics.Extents(font, size, msg)

	// create canvas
	ximg := xgraphics.New(X, image.Rect(pos_x, pos_y, 600, 400))
	ximg.For(func(x, y int) xgraphics.BGRA {
		return bg
	})

	// write the text
	_, _, err = ximg.Text(pos_x+2, pos_y+2, fg, size, font, msg)
	checkError(err)

	//win := ximg.XShow()

	msg2 := "blahbéhdweohde"
	_, _, err = ximg.Text(pos_x+2, pos_y+height, fg, size, font, msg2)
	checkError(err)

	secw, sech := xgraphics.Extents(font, size, msg2)

	bounds := image.Rect(pos_x, pos_y+height, pos_x+secw, pos_x+height+sech)

	// get xproto.Window
	winID, err := xproto.NewWindowId(X.Conn())
	checkError(err)

	img := ximg.SubImage(bounds)
	if img == nil {
		fmt.Println("ximg.SubImage is nil")
		os.Exit(1)
	}

	err = ximg.XSurfaceSet(winID)
	checkError(err)

	err = img.XDrawChecked()
	checkError(err)

	win := ximg.XShow()
	ximg.XPaint(win.Id)
}
