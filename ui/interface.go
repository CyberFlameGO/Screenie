package ui

import (
	fyne "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var (
	app = fyne.New()
)

func InitInterface() {
	w := app.NewWindow("Screenie")
	//w.SetFixedSize(true)

	w.SetContent(widget.NewVBox(widget.NewLabel("Init"),
		widget.NewButton("Quit", func() {
			app.Quit()
		})))

	w.ShowAndRun()
}
