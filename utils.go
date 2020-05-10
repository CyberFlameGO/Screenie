package Screenie

import (
	"github.com/lxn/win"
)

var hDC = win.GetDC(0) //todo change this to the screen that is active

func GetScreenHeight() int {
	return int(win.GetDeviceCaps(hDC, win.HORZRES))
}

func GetScreenWidth() int {
	return int(win.GetDeviceCaps(hDC, win.VERTRES))
}
