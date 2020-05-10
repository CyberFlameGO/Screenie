package ui

import (
	"fmt"
	fyne "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"log"
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
			err := capture.SaveScreenshot(fmt.Sprintf("MM-DD-YYYY", time.Now())) //todo ask for them to put in the name
			if err != nil {
				log.Fatalf("could not start new recording at: %s with error: %v", time.Now(), err.Error())
			}
		}),
		widget.NewButton("Stop recording", func() {
			capture.StopRecording()
		})))

	w.ShowAndRun()
}
