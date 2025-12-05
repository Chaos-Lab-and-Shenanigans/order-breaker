package rickroll

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
)

func RestoreDesktop() {
	config.GotRickRolled = false
	errCh1 := make(chan error)
	errCh2 := make(chan error)

	go restoreWallpaper(errCh1)
	go restoreNames(errCh2)
	go stopAudio()

	err := <-errCh1
	if err != nil {
		config.Cfg.LogsCh <- fmt.Sprintf("Error restoring wallpaper: %v", err)
		return
	}

	err = <-errCh2
	if err != nil {
		config.Cfg.LogsCh <- fmt.Sprintf("Error occured while restoring file: %v", err)
		return
	}

	setWindowRestore()
	config.Cfg.LogsCh <- "Successfully recovered the madness"
}

func setWindowRestore() {
	descL := CenteredLabel("Restored desktop successfully!")
	descL.TextStyle.Bold = true

	if config.IsSlideshowWall {
		descL.SetText(descL.Text + "\n" + "Change your wallpaper yourself")
	}

	quote := CenteredLabel("\"Why suffer alone when you have friends\"")
	quote.TextStyle.Italic = true

	config.Cfg.Window.SetCloseIntercept(func() { config.Cfg.ControlCh <- "exit" })

	config.Cfg.Window.SetContent(container.NewVBox(
		descL,
		CenteredLabel("Share with your single friends."),
		quote,
		CenteredLabel("Those are some wise words indeed.\nIf you know what I mean."),
		CenteredLabel("🙂"),
		layout.NewSpacer(),
		widget.NewSeparator(),
		config.GetHomeExitButtons(),
	),
	)
}

func CenteredLabel(s string) *widget.Label {
	label := widget.NewLabel(s)
	label.Alignment = fyne.TextAlignCenter
	return label
}
