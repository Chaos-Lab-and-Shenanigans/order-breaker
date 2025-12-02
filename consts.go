package main

import (
	_ "embed"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

//go:embed assets/wall.png
var rickWall []byte

//go:embed assets/audio.mp3
var rickAudio []byte

var (
	windowSize    = fyne.NewSize(300, 300)
	HOMEDIR, _    = os.UserHomeDir()
	PATH_DESKTOP  = filepath.Join(HOMEDIR, "Desktop")
	PATH_BACKUPOB = filepath.Join(PATH_DESKTOP, ".backupOB")
	DATABASE      = "backupob.db"
)

// Main window's
var (
	myApp     = app.New()
	windowAst = myApp.NewWindow("Astrology")
	logsL     = widget.NewLabel("")
	logsCh    = setUpdaterChannel(logsL)
	restartCh = getRestartChannel()
)
