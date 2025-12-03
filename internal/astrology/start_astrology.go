package astrology

import (
	"fyne.io/fyne/v2"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
)

func StartAstro() func() {
	return func() {
		firstPage()
	}
}

func InitConfig(path string, dbName string, rickAudioBytes *[]byte, rickWallBytes *[]byte, w fyne.Window, logsCh chan string, controlCh chan string) {
	config.Cfg.LogsCh = logsCh
	config.Cfg.ControlCh = controlCh
	config.Cfg.Path = path
	config.Cfg.DBName = dbName
	config.Cfg.RickyAudioBytes = rickAudioBytes
	config.Cfg.RickyWall = rickWallBytes
	config.Cfg.Window = w
}
