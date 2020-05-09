package capture

import (
	"errors"
	"github.com/cretz/go-scrap"
	"image/png"
	"os"
	"time"
)

func SaveScreenshot(fileName string) error {
	img, err := getScreenshot()
	if err != nil {
		return errors.New("could not find screenshot, with error: " + err.Error())
	}

	file, err := os.Create(fileName)

	if err != nil {
		return errors.New("could not create file, with error: " + err.Error())
	}

	defer file.Close()
	return png.Encode(file, img)
}

func getScreenshot() (*scrap.FrameImage, error) {
	//make sure that the dpi is aware
	if err := scrap.MakeDPIAware(); err != nil {
		return nil, errors.New("could not make DPI aware, with error: " + err.Error())
	}

	//get the primary display
	//TODO change this from primary display to whichever one
	//is actually selected
	d, err := scrap.PrimaryDisplay()
	if err != nil {
		return nil, errors.New("could not get the primary display, with error: " + err.Error())
	}

	c, err := scrap.NewCapturer(d)
	if err != nil {
		return nil, err
	}

	for {
		if img, _, err := c.FrameImage(); img != nil || err != nil {
			//detech so that the method is safe to use after
			if img != nil {
				img.Detach()
			}

			return img, errors.New("error with image detaching, with error: " + err.Error())
		}

		time.Sleep(17 * time.Millisecond)
	}
}
