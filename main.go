package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/astrology"
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/sqlite3"
)

// Fix bug: Store backup wall path on db
func main() {
	astrology.InitConfig(&rickAudio, &rickWall, windowAst, logsCh, controlCh)
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
		logContent := container.NewScroll(logsL)
		// Set a reasonable minimum size for the log content area
		logContent.SetMinSize(config.LogsWindowSize)

		logWindow.SetContent(container.NewVBox(
			logContent,
		))

		if showAndRun {
			logWindow.ShowAndRun()
		} else {
			logWindow.Show()
		}
	}
}

func setStartWindow() {
	welcomeL := widget.NewLabel("Welcome to astrology!\nEnter your details and let the app guess your personality ")
	windowAst.SetContent(container.NewVBox(
		welcomeL,
		config.Extra,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewButton("Start Astrology", astrology.StartAstro()),
		widget.NewButton("Compatibility checker", astrology.StartCompatibilityChecker()),
		widget.NewButton("Exit", config.HandleExit),
	))

	windowAst.Resize(config.AstWindowSize)
}
