package main

import(
	"github.com/BurntSushi/xgbutil/xgraphics"
	"image"
	"github.com/BurntSushi/xgbutil"
	"os"
	"fmt"

)

func drawMenu(X *xgbutil.XUtil, pos_x, pos_y int, size float64) {
	// background color of the canvas
	var bg = xgraphics.BGRA{B: 0xff, G: 0x66, R: 0x33, A: 0xff}

	// color of the text
	var fg = xgraphics.BGRA{B: 0xff, G: 0xff, R: 0xff, A: 0xff}

	// select font
	fontReader, err := os.Open("/usr/share/fonts/dejavu/DejaVuSansMono.ttf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fontReader.Close()

	// parse font
	font, err := xgraphics.ParseFont(fontReader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	msg := "it works"
	// get proper width and height of the 1 line text
	width, height := xgraphics.Extents(font, size, msg)

	// create canvas
	ximg := xgraphics.New(X, image.Rect(pos_x, pos_y, width, height))
	ximg.For(func(x, y int) xgraphics.BGRA {
		return bg
	})

	// write the text
	_, _, err = ximg.Text(pos_x, pos_y, fg, size, font, msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	win := ximg.XShow()

	msg2 := "blahbéhdweohde"
	_, _, err = ximg.Text(pos_x, pos_y+height, fg, size, font, msg2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	secw, sech := xgraphics.Extents(font, size, msg2)

	bounds := image.Rect(pos_x, pos_y+height, secw+width, sech+height+sech)

	// get xproto.Window
//	winID, err := xproto.NewWindowId(X.Conn())
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}

	img := ximg.SubImage(bounds)

//	err = img.XSurfaceSet(winID)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}

	err = img.XDrawChecked()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ximg.XPaint(win.Id)
}
