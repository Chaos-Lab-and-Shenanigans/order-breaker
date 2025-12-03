package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	astrology "github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/astrology"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/sqlite3"
)

// Fix bug: Store backup wall path on db
func main() {
	astrology.InitConfig(PATH_BACKUPOB, DATABASE, &rickAudio, &rickWall, windowAst, logsCh, controlCh)
	db, err := sqlite3.CreateAndConnect()
	if err != nil { //Critical error, only show logs
		er := fmt.Sprintf("Error connecting to database: %v", err)
		logsCh <- er
		fyne.Do(showLogs(true))
		return
	}
	defer db.Close()

	setStartWindow()

	fyne.Do(showLogs(false))
	windowAst.ShowAndRun()
}

func showLogs(showAndRun bool) func() {
	return func() {
		logWindow := myApp.NewWindow("Logs")

		// 1. Wrap the dynamic log content in a scroll container
		logContent := container.NewVScroll(logsL)
		// Set a reasonable minimum size for the log content area
		logContent.SetMinSize(fyne.NewSize(250, 150))

		logWindow.SetContent(container.NewVBox(
			logContent,
		))

		if showAndRun {
			logWindow.ShowAndRun()
		}

		logWindow.Show()
	}
}

func setStartWindow() {
	welcomeL := widget.NewLabel("Welcome to astrology!\nEnter your details and let the app guess your personality ")
	windowAst.SetContent(container.NewVBox(
		welcomeL,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewButton("Start Astrology", astrology.StartAstro()),
		widget.NewButton("Compatibility checker", astrology.StartCompatibilityChecker()),
		widget.NewButton("Exit", config.HandleExit),
		widget.NewSeparator(),
	))

	windowAst.Resize(windowSize)
}
