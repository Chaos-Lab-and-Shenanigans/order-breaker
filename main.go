package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	astrology "github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/astrology"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/sqlite3"
	tappedfunctions "github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/tapped_functions"
)

func main() {
	db, err := sqlite3.CreateAndConnect(PATH_BACKUPOB, PATH_DB, logsCh)
	if err != nil {
		er := fmt.Sprintf("Error connecting to database: %v", err)
		windowAst.SetContent(tappedfunctions.CenteredLabel(er + "\n" + "Closing in 5 seconds..."))
		time.Sleep(5 * time.Second)
		return
	}
	defer db.Close()

	astrology.InitConfig(db, PATH_BACKUPOB, PATH_DB, &rickAudio, &rickWall, windowAst, logsCh, restartCh)

	setStartWindow()

	go showLogs()
	windowAst.ShowAndRun()
}

func showLogs() {
	logWindow := myApp.NewWindow("Logs")

	// 1. Wrap the dynamic log content in a scroll container
	logContent := container.NewVScroll(logsL)
	// Set a reasonable minimum size for the log content area
	logContent.SetMinSize(fyne.NewSize(250, 150))

	logWindow.SetContent(container.NewVBox(
		logContent,
	))

	logWindow.Show()
}
func setStartWindow() {
	windowAst.SetContent(container.NewVBox(
		welcomeL,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewButton("Start Astrology", astrology.StartAstro()),
		widget.NewButton("Compatibility checker", astrology.StartCC()),
		widget.NewButton("Exit", func() { myApp.Quit() }),
		widget.NewSeparator(),
	))

	windowAst.Resize(windowSize)
}
