package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	handleastrology "github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/astrology"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/sqlite3"
)

func main() {
	app := app.New()
	window := app.NewWindow("Main")
	backupL := widget.NewLabel("Logs will appear here")
	logsCh := SetUpdaterChannel(backupL)

	db, err := sqlite3.CreateAndConnect(logsCh, BACKUPOB, DATABASE)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()

	window.SetContent(container.NewVBox(
		welcomeL,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewButton("Start Astrology?", handleastrology.StartAstro(window, logsCh)),
		widget.NewButton("Exit", func() { app.Quit() }),
		widget.NewSeparator(),
		backupL,
	))

	window.Resize(windowSize)
	window.ShowAndRun()
}
