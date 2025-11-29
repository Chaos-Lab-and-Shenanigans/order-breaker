package main

import (
	_ "embed"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//go:embed assets/wall.png
var picByte []byte

//go:embed assets/audio.mp3
var rickAudio []byte

var (
	windowSize = fyne.NewSize(300, 300)
	welcomeL   = widget.NewLabel("Welcome to astrology!\nEnter your details and let the app guess your personality ")
	HOMEDIR, _ = os.UserHomeDir()
	DESKTOP    = filepath.Join(HOMEDIR, "Desktop")
	BACKUPOB   = filepath.Join(DESKTOP, ".backupOB")
	DATABASE   = filepath.Join(BACKUPOB, "backupob.db")
)
