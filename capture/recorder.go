package capture

import (
	"errors"
	"github.com/cretz/go-scrap"
)

//generates a new capturer from the scrap library
//the primary usage of this is to ensure that the
//displays are not null and that the capturer is being
//set up correctly
func capturer() (*scrap.Capturer, error) {
	//ensure that the display that is being captured
	//is not null and wont be causing issues
	//TODO change this from primary display to whichever one
	//is actually selected
	if d, err := scrap.PrimaryDisplay(); err != nil {
		return nil, errors.New("could not link capturer to display, with error: " + err.Error())
	} else if c, err := scrap.NewCapturer(d); err != nil {
		return nil, errors.New("could not create new capturer, with error: " + err.Error())
	} else {
		return c, nil
	}
}
