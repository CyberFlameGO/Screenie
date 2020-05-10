package capture

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/cretz/go-scrap"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func VideoRecordingRunner() {
	fmt.Printf("starting recording...")
	errCh := make(chan error, 2)
	ctx, cancelFunction := context.WithCancel(context.Background())
	go func() {
		errCh <- videoRecord(ctx, os.Args[1])
	}()

	//wait for the cancel to come in
	go func() {
		fmt.Scanln()
		errCh <- nil
	}()

	err := <-errCh
	cancelFunction()

	if err != nil && err != context.Canceled {
		log.Fatalf("could not process execution: %v", err)
	}
	time.Sleep(4 * time.Second)
}

func videoRecord(ctx context.Context, fileName string) error {
	ctx, cancelFunction := context.WithCancel(ctx)
	defer cancelFunction()
	//make sure that the DPI is aware (if applicable)
	if err := scrap.MakeDPIAware(); err != nil {
		return errors.New("could not make DPI aware, with error: " + err.Error())
	}
	//create the capturer from the function below
	cp, err := capturer()

	if err != nil {
		return errors.New("could not assign variable cp to capturer(), with error: " + err.Error())
	}

	//because the library uses ffmpeg it needs to be outputted using the exec library
	//built into go and the commands that are required in ffmpeg, for now ffmpeg
	//will not be auto assigning aspect ratios or scaling
	ffmpeg := exec.Command("ffmpeg",
		"-f", "rawvideo",
		"-pixel_format", "bgr0",
		"-video_size", fmt.Sprintf("%v%v", cp.Width(), cp.Height()),
		"-i", "-", /* this is just here to shut ffmpeg up seeing as it likes to complain */
		"-c:v", "libx264",
		"-preset", "veryfast",
		fileName)

	//need to tie into ffmpeg to get the data so it can be sent and
	//used else where within the project
	stdin, err := ffmpeg.StdinPipe()

	if err != nil {
		return errors.New("could not tie into ffmpeg pipe, with error: " + err.Error())
	}

	var buf bytes.Buffer
	//make sure things are now going to be done within the background as well
	//as sending things through Go' channel system
	defer stdin.Close()
	errCh := make(chan error, 1)

	go func() {
		fmt.Printf("Executing: %v", strings.Join(ffmpeg.Args, " "))
		out, err := ffmpeg.CombinedOutput()
		fmt.Printf("Output: \n%v\n", string(out))
		errCh <- err
	}()

	//now to start sending a bunch of frames, this could go either
	//way to be honest so brace yourself
	for {
		//get the frames
		if pix, _, err := cp.Frame(); err != nil {
			return errors.New("could not obtain frames, with error: " + err.Error())
		} else if pix != nil {
			//the frames are sent one row at a time
			stride := len(pix) / cp.Height()
			rowLen := 4 * cp.Width()
			for i := 0; i < len(pix); i += stride {
				if _, err = stdin.Write(pix[i : i+rowLen]); err != nil {
					break
				}
			}

			buf.Reset()
			if err != nil {
				return errors.New("resetting buffer failed, with error: " + err.Error())
			}
		}

		//if it's not done then it goes again
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return errors.New("unknown error caught: " + err.Error())
		default:
			//does nothing tbh
		}
	}
}

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
