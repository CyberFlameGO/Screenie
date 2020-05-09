package capture

import (
	"errors"
	"github.com/cretz/go-scrap"
	"time"
)

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
