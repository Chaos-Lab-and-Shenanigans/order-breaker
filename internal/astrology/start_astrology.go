package astrology

import (
	"database/sql"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
)

func StartAstro() func() {
	return func() {
		initNav()
		firstPage()
	}
}

func initNav() {
	homeB := widget.NewButton("Home", func() { sendRestart(config.Cfg.RestartCh) })
	exitB := widget.NewButton("Exit", func() { fyne.CurrentApp().Quit() })

	config.NavigationButtons = container.New(layout.NewGridLayout(2), exitB, homeB)
}

func sendRestart(ch chan string) {
	select {
	case ch <- "restart":
	default:
		time.Sleep(time.Millisecond)
	}
}

func InitConfig(db *sql.DB, path string, dbName string, rickAudioBytes *[]byte, rickWallBytes *[]byte, w fyne.Window, logsCh chan string, restartCh chan string) {
	config.Cfg.DB = db
	config.Cfg.LogsCh = logsCh
	config.Cfg.RestartCh = restartCh
	config.Cfg.Path = path
	config.Cfg.PathDB = filepath.Join(path, dbName)
	config.Cfg.RickyAudioBytes = rickAudioBytes
	config.Cfg.RickyWall = rickWallBytes
	config.Cfg.Window = w
}
