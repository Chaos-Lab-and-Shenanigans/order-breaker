package main

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed assets/wall.png
var picByte []byte

var (
	HOMEDIR, _ = os.UserHomeDir()
	DESKTOP    = filepath.Join(HOMEDIR, "Desktop")
	BACKUPOB   = filepath.Join(HOMEDIR, ".backupOB")
	DATABASE   = filepath.Join(BACKUPOB, "backupob.db")
)
