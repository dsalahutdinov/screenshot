package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)

var ops uint64

func screenshot(targetUrl string, width string, height string) (string, error) {
	screenshotPath := fmt.Sprintf("/tmp/screenshot%d.png", atomic.AddUint64(&ops, 1))

	fmt.Println(screenshotPath)
	var chromeArguments = []string{
		"--headless",
		"--disable-gpu",
		"--hide-scrollbars",
		"--disable-crash-reporter",
		"--screenshot=" + screenshotPath,
    "--window-size=" + width + "," + height,
	}
	if os.Geteuid() == 0 {
		chromeArguments = append(chromeArguments, "--no-sandbox")
	}
	chromeArguments = append(chromeArguments, targetUrl)

	// get a context to run the command in
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "google-chrome-stable", chromeArguments...)
	cmd.Env = os.Environ()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Wait for the screenshot to finish and handle the error that may occur.
	if err := cmd.Wait(); err != nil {

		// If if this error was as a result of a timeout
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Timeout %s", err)
			return "", err
		}

		fmt.Println("Screenshoot failed")

		return "", err
	}
	return screenshotPath, nil
}

func screenshotHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
  width:= r.URL.Query().Get("width")
  height := r.URL.Query().Get("height")
	screenshot, _ := screenshot(url, width, height)
	http.ServeFile(w, r, screenshot)
	os.RemoveAll(screenshot)
}

func main() {
	http.HandleFunc("/screenshot", screenshotHandler)
	http.ListenAndServe(":8080", nil)
}
