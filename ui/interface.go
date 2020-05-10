package ui

import (
	"fmt"
	fyne "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"screenie/capture"
	"time"
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
		}),
		widget.NewButton("Record", func() {
			capture.VideoRecordingRunner()
		}),
		widget.NewButton("Screenshot", func() {
			capture.SaveScreenshot(fmt.Sprintf("MM-DD-YYYY", time.Now())) //todo ask for them to put in the name
		})))

	w.ShowAndRun()
}
