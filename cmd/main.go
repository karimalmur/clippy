package main

import (
	"context"
	"regexp"

	"github.com/getlantern/systray"
)

var (
	cancelWatch *context.CancelFunc
	urlRegex    *regexp.Regexp
)

func main() {
	urlRegex = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
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
