package main

import (
	"os"
	"path/filepath"
)

var (
	HOMEDIR, _ = os.UserHomeDir()
	DESKTOP    = filepath.Join(HOMEDIR, "Desktop")
	BACKUPOB   = filepath.Join(HOMEDIR, ".backupOB")
	DATABASE   = filepath.Join(BACKUPOB, "backupob.db")
)
