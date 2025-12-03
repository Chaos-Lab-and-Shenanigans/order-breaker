package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func setUpdaterChannel(x *widget.Label) chan string {
	const logBatchSize = 10
	channel := make(chan string)
	x.Selectable = true

	go func() {
		// Counter to track how many messages have been received in the current batch
		messageCount := 0

		for text := range channel {
			messageCount++
			fyne.Do(func() {
				// 2. Check if the counter hits the batch size threshold
				if messageCount > 0 && messageCount%logBatchSize == 0 {
					// Insert an extra blank line to separate the batches
					x.SetText(x.Text + "\n\n" + text)
				} else {
					// Regular append
					x.SetText(x.Text + "\n" + text)
				}
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
