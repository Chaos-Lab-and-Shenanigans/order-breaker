package tappedfunctions

import (
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RestoreDesktop(db *sql.DB, path string, pathDB string, w fyne.Window, logsCh chan string, restartCh chan string) {
	errCh1 := make(chan error)
	errCh2 := make(chan error)

	go restoreWallpaper(errCh1)
	go restoreNames(db, path, errCh2)
	go stopAudio()

	err := <-errCh1
	if err != nil {
		logsCh <- fmt.Sprintf("%v", err)
		return
	}

	err = <-errCh2
	if err != nil {
		logsCh <- fmt.Sprintf("Successfully recovered the madness\n %v", err)
		return
	}

	setWindowRestore(w, restartCh)

	logsCh <- "Successfully recovered the madness"
}

func setWindowRestore(w fyne.Window, restartCh chan string) {
	descL := CenteredLabel("Restored desktop successfully!")
	descL.TextStyle.Bold = true

	quote := CenteredLabel("\"Why suffer alone when you have friends\"")
	quote.TextStyle.Italic = true

	navigation := container.New(
		layout.NewGridLayout(2),
		widget.NewButton("Start again", func() { restartCh <- "restart" }),
		widget.NewButton("Exit", func() { fyne.CurrentApp().Quit() }),
	)

	w.SetContent(container.NewVBox(
		descL,
		CenteredLabel("Share with your single friends."),
		quote,
		CenteredLabel("Those are some wise words indeed.\nIf you know what I mean."),
		CenteredLabel("🙂"),
		layout.NewSpacer(),
		widget.NewSeparator(),
		navigation,
	),
	)
}

func CenteredLabel(s string) *widget.Label {
	label := widget.NewLabel(s)
	label.Alignment = fyne.TextAlignCenter
	return label
}
