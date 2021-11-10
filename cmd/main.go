package main

import (
	"context"

	"github.com/getlantern/systray"
)

var (
	cancelWatch *context.CancelFunc
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	SetupTray()
	WatchClipboard()
}

func onExit() {
	if cancelWatch != nil {
		(*cancelWatch)()
	}

	if pauseTimer != nil {
		pauseTimer.Stop()
	}
}
