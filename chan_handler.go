package main

import (
	"time"

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
				// Check if the counter hits the batch size threshold
				if messageCount > 0 && messageCount%logBatchSize == 0 {
					// Insert an extra blank line to separate the batches
					x.SetText(x.Text + "\n\n" + text)
				} else {
					x.SetText(x.Text + "\n" + text)
				}
			})
		}
	}()
	return channel
}

func getControlCh() chan string {
	controlCh := make(chan string)
	go func() {
		for msg := range controlCh {
			fyne.Do(respond(msg))
		}
	}()

	return controlCh
}

func respond(msg string) func() {
	return func() {
		switch msg {
		case "restart":
			fyne.Do(func() { setStartWindow() })
		case "ghost":
			fyne.Do(ghostUser)
		case "exit":
			fyne.Do(exit)
		}
	}
}

func exit() {
	fyne.CurrentApp().Quit()
}

// Show the screen again and again
func ghostUser() {
	go func() {
		i := 0
		for {
			if i == 0 { //Go to home page if Home button was pressed first time
				fyne.Do(func() { setStartWindow() })
				i = 1
			} else { //Ghost otherwise
				fyne.Do(func() { windowAst.Show() })
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
