package main

import (
	"context"

	"golang.design/x/clipboard"
)

func WatchClipboard() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := clipboard.Watch(ctx, clipboard.FmtText)
	cancelWatch = &cancel

	go func() {
		for data := range ch {
			println(string(data))
		}
	}()
}

func TryCancelWatch() {
	if cancelWatch == nil {
		return
	}

	(*cancelWatch)()
}
