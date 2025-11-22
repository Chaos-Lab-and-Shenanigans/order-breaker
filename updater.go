package main

import "fyne.io/fyne/v2"

type SetText interface {
	SetText(string)
}

func SetUpdaterChannel(x SetText) chan string {
	channel := make(chan string)
	go func() {
		for text := range channel {
			fyne.Do(func() {
				x.SetText(text)
			})
		}
	}()
	return channel
}
