package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/sqlite3"
	tappedfunctions "github.com/Chaos-Lab-and-Shenanigans/order-breaker/tapped_functions"
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
		backupL,
		widget.NewButton("Restore desktop", tappedfunctions.TapRestoreDesktop(db, logsCh, BACKUPOB, DATABASE)),
		widget.NewButton("Ricky", tappedfunctions.TapRickRollDesktop(db, logsCh, BACKUPOB, picByte)),
		widget.NewButton("Exit", func() { app.Quit() }),
	))

	window.ShowAndRun()
}
