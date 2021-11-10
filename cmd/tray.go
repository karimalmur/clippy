package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/getlantern/systray"
)

const (
	pauseDuration   = math.MaxInt16
	pauseDuration5  = time.Minute * 5
	pauseDuration10 = time.Minute * 10
)

var pauseTimer *time.Timer

func SetupTray() {
	icon, err := ioutil.ReadFile("./assets/icon.ico")
	if err != nil {
		fmt.Print("Failed to open icon file", err)
	}

	systray.SetIcon(icon)
	systray.SetTitle("Clippy")
	systray.SetTooltip("Clipboard url logger")

	enable := systray.AddMenuItem("Enable", "Enable")
	enable.Disable()

	pause := systray.AddMenuItem("Pause", "Pause")
	pauseind := pause.AddSubMenuItem("Indefinitely", "Pause Indefinitely")
	pause5 := pause.AddSubMenuItem("5 Mins", "Pause for 5 mins")
	pause10 := pause.AddSubMenuItem("10 Mins", "Pause for 10 mins")

	quit := systray.AddMenuItem("Quit", "Quit")

	go func() {
		for {
			select {
			case <-enable.ClickedCh:
				togglePause(pause, enable, 0)
				WatchClipboard()
			case <-pauseind.ClickedCh:
				togglePause(pause, enable, pauseDuration)
				TryCancelWatch()
			case <-pause5.ClickedCh:
				togglePause(pause, enable, pauseDuration5)
				TryCancelWatch()
			case <-pause10.ClickedCh:
				togglePause(pause, enable, pauseDuration10)
				TryCancelWatch()
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func togglePause(pauseMenuItem *systray.MenuItem, resumeMenuItem *systray.MenuItem, pauseDurationMinutes time.Duration) {
	if pauseTimer != nil {
		pauseTimer.Stop()
		pauseTimer = nil
	}

	if pauseMenuItem.Disabled() {
		pauseMenuItem.Enable()
		resumeMenuItem.Disable()
	} else {
		pauseMenuItem.Disable()
		resumeMenuItem.Enable()

		if pauseDurationMinutes < math.MaxInt16 {
			timer := time.NewTimer(time.Minute * pauseDurationMinutes)
			pauseTimer = timer

			go func() {
				<-timer.C
				WatchClipboard()
			}()
		}
	}
}
