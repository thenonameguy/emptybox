package main

import (
	"fmt"
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
	msg2 := "jfewőjdeiodhjwedoeh"
	secw, sech := xgraphics.Extents(font, size, msg2)

	// create canvas(x resource pixmap)
	ximg := xgraphics.New(X, image.Rect(pos_x, pos_y, pos_x+secw, pos_y+height+sech))
	ximg.For(func(x, y int) xgraphics.BGRA {
		return bg
	})

	// write the text
	_, _, err = ximg.Text(pos_x, pos_y, fg, size, font, msg)
	checkError(err)

	_, _, err = ximg.Text(pos_x, pos_y+height, fg, size, font, msg2)
	checkError(err)

	// now update where we have written text
	bounds := image.Rect(pos_x, pos_y+height, pos_x+secw, pos_y+height+sech)

	img := ximg.SubImage(bounds)
	if img == nil {
		fmt.Println("ximg.SubImage is nil")
		os.Exit(1)
	}

	img.XDraw()

	win := ximg.XShow()
	ximg.XPaint(win.Id)
}
