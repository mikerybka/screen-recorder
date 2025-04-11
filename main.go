package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	for {
		err := recordUntilHourTick()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func recordUntilHourTick() error {
	now := time.Now()
	outfile := fmt.Sprintf("/home/mike/screen-recordings/%04d-%02d-%02d/%d.mkv", now.Year(), now.Month(), now.Day(), now.Hour())
	err := os.MkdirAll(filepath.Dir(outfile), os.ModePerm)
	if err != nil {
		return err
	}
	d := 3600 - (now.Minute()*60 + now.Second())
	cmd := exec.Command("ffmpeg",
		"-video_size", "3840x2160",
		"-framerate", "30",
		"-f", "x11grab",
		"-i", ":0.0",
		"-d", strconv.Itoa(d),
		"-y",
		outfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
