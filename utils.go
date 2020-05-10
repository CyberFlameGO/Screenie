package Screenie

import (
	"fyne.io/fyne/widget"
	"github.com/lxn/win"
)

var hDC = win.GetDC(0) //todo change this to the screen that is active

func GetScreenHeight() int {
	return int(win.GetDeviceCaps(hDC, win.HORZRES))
}

func GetScreenWidth() int {
	return int(win.GetDeviceCaps(hDC, win.VERTRES))
}

func FyneLabel(text string) *widget.Label {
	w := widget.NewLabel(text)
	return w
}

func FyneEntry(placeHolder string) *widget.Entry {
	w := widget.NewEntry()
	w.SetPlaceHolder(placeHolder)
	return w
}

func FyneButton(text string, action func()) *widget.Button {
	w := widget.NewButton(text, action)
	return w
}
