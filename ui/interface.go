package ui

import (
	"fmt"
	"fyne.io/fyne"
	f "fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"log"
	Screenie "screenie"
	"screenie/capture"
	"time"
)

var (
	app      = f.New()
	settings Settings
)

type Settings struct {
	RecordingScreen string
	Theme           string
}

func InitInterface() {
	app.Settings().SetTheme(theme.DarkTheme()) //todo get the theme from the config
	w := app.NewWindow("Screenie")
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewFormLayout(),
		Screenie.FyneGroup("primary_interface", true,
			fyne.NewContainerWithLayout(layout.NewGridLayout(1),
				Screenie.FyneButton("Quit", func() {
					app.Quit()
				}),
				fyne.NewContainerWithLayout(layout.NewGridLayout(1),
					Screenie.FyneButton("Start Recording", func() {
						capture.VideoRecordingRunner()
					}),
					fyne.NewContainerWithLayout(layout.NewGridLayout(1),
						Screenie.FyneButton("Stop Recording", func() {
							capture.StopRecording()
						}),
						fyne.NewContainerWithLayout(layout.NewGridLayout(1),
							Screenie.FyneButton("Screenshot", func() {
								err := capture.SaveScreenshot(fmt.Sprintf("DD-MM-YYYY", time.Now()))
								//todo add this into a error box
								if err != nil {
									log.Fatalf("could not save screenshot at: %s with error: %v", time.Now(), err.Error())
								}
							}),
							fyne.NewContainerWithLayout(layout.NewGridLayout(1),
								Screenie.FyneButton("Settings", func() {
									settingsWindow()
								})))))))))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()

}

func settingsWindow() {

	settings.RecordingScreen = Screenie.FyneEntry(string(1)).Text
	settings.Theme = Screenie.FyneEntry("Dark").Text
	//todo set the theme depending on what is selected
	w := app.NewWindow("Screenie settings")
	w.SetContent(fyne.NewContainerWithLayout(layout.NewFormLayout(),
		Screenie.FyneLabel("Screen to record"),
		Screenie.FyneEntry(string(1))))

	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}
