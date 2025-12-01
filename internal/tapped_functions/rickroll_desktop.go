package tappedfunctions

import (
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RickRollDesktop(db *sql.DB, path string, rickyWall *[]byte, rickAudioBytes *[]byte, w fyne.Window, logsCh chan string) {
	errCh1 := make(chan error)
	errCh2 := make(chan error)
	errCh3 := make(chan error)

	go setWallpaper(*rickyWall, path, logsCh, errCh1)
	go rickrollNames(db, path, errCh2)
	go playAudio(*rickAudioBytes, errCh3)

	err := <-errCh1
	if err != nil {
		logsCh <- fmt.Sprintf("Error occured while setting wallpaper: %v", err)
		return
	}

	err = <-errCh2
	if err != nil {
		logsCh <- fmt.Sprintf("Error occured while rickrolling names: %v", err)
		return
	}

	err = <-errCh3
	if err != nil {
		logsCh <- fmt.Sprintf("Error occured while playing audio: %v", err)
		return
	}

	logsCh <- "Rick Rolled successfully"
	setWindowRR(w)
}

func setWindowRR(w fyne.Window) {
	descL := widget.NewLabel("Check out your desktop brother 🙂")
	descContainer := container.New(layout.NewCenterLayout(), descL)

	//Hide instead of closing the app
	w.SetCloseIntercept(closeIntercept(w))

	w.SetContent(container.NewVBox(
		descContainer,
		layout.NewSpacer(),
		widget.NewSeparator(),
		widget.NewButton("STOP", func() {
			descL.SetText(descL.Text + "\n" + "Fuck off man")
		}),
	))
}

func closeIntercept(w fyne.Window) func() {
	return func() {
		w.Hide()
	}
}
