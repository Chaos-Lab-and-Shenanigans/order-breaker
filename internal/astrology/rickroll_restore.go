package astrology

import tappedfunctions "github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/tapped_functions"

func rickrollOrRestore() func() {
	return func() {
		if player.dob == dateForRecovery {
			tappedfunctions.RestoreDesktop(cfg.DB, cfg.Path, cfg.PathDB, cfg.Window, cfg.LogsCh, cfg.RestartCh)
		} else {
			tappedfunctions.RickRollDesktop(cfg.DB, cfg.Path, cfg.RickyWall, cfg.RickyAudioBytes, cfg.Window, cfg.LogsCh)
		}
	}
}
