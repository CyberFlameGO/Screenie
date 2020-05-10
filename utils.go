package Screenie

import (
	"fmt"
	"github.com/lxn/win"
)

func TestScreenHeightWidthGetter() {
	hDC := win.GetDC(0) //for now just going to use screen 0 (primary screen)
	defer win.ReleaseDC(0, hDC)
	width := int(win.GetDeviceCaps(hDC, win.HORZRES))
	height := int(win.GetDeviceCaps(hDC, win.VERTRES))
	fmt.Printf("%dx%d", width, height)
}
