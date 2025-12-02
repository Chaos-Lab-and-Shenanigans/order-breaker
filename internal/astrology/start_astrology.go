package astrology

import (
	"database/sql"
	"path/filepath"

	"fyne.io/fyne/v2"
)

func StartAstro() func() {
	return func() {
		firstPage()
	}
}

func InitConfig(db *sql.DB, path string, dbName string, rickAudioBytes *[]byte, rickWallBytes *[]byte, w fyne.Window, logsCh chan string, restartCh chan string) {
	cfg.DB = db
	cfg.LogsCh = logsCh
	cfg.RestartCh = restartCh
	cfg.Path = path
	cfg.PathDB = filepath.Join(path, dbName)
	cfg.RickyAudioBytes = rickAudioBytes
	cfg.RickyWall = rickWallBytes
	cfg.Window = w
}
