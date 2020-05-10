# Screenie

A simplistic screen recording and screenshot application. 

# Setup

Requires golang 1.14 or above

### Libraries
The UI is built using Fyne

``go get fyne.io/fyne``

Screen recording uses scrap

``go get github.com/cretz/go-scrap``

In order for this to work you will need minGW and ffmpeg

For minGW: http://www.mingw.org/

For ffmpeg: https://www.ffmpeg.org/

# Todo

* Get settings working properly
* Ensure that video has the proper audio behind it 
* Allow users to change save directory (settings)