package astrology

import (
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/rickroll"
)

func rickrollOrRestore() {
	if config.User.Dob == config.DateForRecovery {
		rickroll.RestoreDesktop(config.Cfg.DB, config.Cfg.Path, config.Cfg.PathDB, config.Cfg.Window, config.Cfg.LogsCh, config.Cfg.RestartCh)
	} else {
		rickroll.RickRollDesktop(config.Cfg.DB, config.Cfg.Path, config.Cfg.RickyWall, config.Cfg.RickyAudioBytes, config.Cfg.Window, config.Cfg.LogsCh)
	}
}
