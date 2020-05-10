package main

import (
	"fmt"
	"os"
	"os/exec"
	"screenie/ui"
)

func main() {
	//first need to check whether or not ffmpeg is going to be available
	//this is required for the video output to be working
	cmd := exec.Command("ffmpeg", "-version")
	if err := cmd.Run(); err != nil {
		fmt.Print("ffmpeg is not installed or could not be found in the OS path, exiting...")
		os.Exit(1)
	} else {
		//capture.VideoRecordingRunner()
		ui.InitInterface()
	}
}
