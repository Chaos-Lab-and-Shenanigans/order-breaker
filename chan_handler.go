package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func setUpdaterChannel(x *widget.Label) chan string {
	channel := make(chan string)
	go func() {
		for text := range channel {
			fyne.Do(func() {
				x.SetText(x.Text + "\n" + text)
			})
		}
	}()
	return channel
}

func getRestartChannel() chan string {
	channel := make(chan string)
	go func() {
		msg := <-channel
		if msg == "restart" {
			fyne.Do(func() { setStartWindow() })
		}
	}()

	return channel
}
