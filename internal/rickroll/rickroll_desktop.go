package rickroll

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
)

func RickRollDesktop() {
	errCh1 := make(chan error)
	errCh2 := make(chan error)
	errCh3 := make(chan error)

	go setWallpaper(errCh1)
	go rickrollNames(errCh2)
	go playAudio(errCh3)

	err := <-errCh1
	if err != nil {
		config.Cfg.LogsCh <- fmt.Sprintf("Error occured while setting wallpaper: %v", err)
		return
	}

	err = <-errCh2
	if err != nil {
		config.Cfg.LogsCh <- fmt.Sprintf("Error occured while rickrolling names: %v", err)
		return
	}

	err = <-errCh3
	if err != nil {
		config.Cfg.LogsCh <- fmt.Sprintf("Error occured while playing audio: %v", err)
		return
	}

	config.Cfg.LogsCh <- "Rick Rolled successfully"
	setWindowRR()
}

func setWindowRR() {
	w := config.Cfg.Window
	descL := widget.NewLabel("Check out your desktop brother 🙂")
	descContainer := container.New(layout.NewCenterLayout(), descL)

	//Hide instead of closing the app
	w.SetCloseIntercept(func() { w.Hide() })

	w.SetContent(container.NewVBox(
		descContainer,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewButton("STOP", func() {
			descL.SetText(descL.Text + "\n" + "Fuck off man")
		}),
	))
}
