package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/kbinani/screenshot"
	log "github.com/sirupsen/logrus"
)

func isValidDisplay(display int) bool {
	n := screenshot.NumActiveDisplays()
	if n < 0 || display >= n {
		return false
	}
	return true
}

func saveFile(filename string, img *image.RGBA) bool {
	file, err := os.Create(filename)
	if err != nil {
		return false
	}
	defer file.Close()
	err = png.Encode(file, img)
	return err == nil
}
func takeScreenshot(display int, filename string) bool {
	if !isValidDisplay(display) {
		log.Error("No display enabled")
		return false
	}
	bounds := screenshot.GetDisplayBounds(display)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Error("Error taking screenshot")
		return false
	}
	if !saveFile(filename, img) {
		log.Error("Error saving file")
		return false
	}
	return true
}

func main() {
	i := 1
	display, _ := strconv.Atoi(os.Args[1])
	seconds, _ := time.ParseDuration(os.Args[2] + "s")
	for {
		filename := fmt.Sprintf("screenshot%d.png", i)
		log.Info("Taking screenshot: ", filename)
		takeScreenshot(display, filename)
		time.Sleep(seconds)
		i++
	}
}
