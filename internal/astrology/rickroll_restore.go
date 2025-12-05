package astrology

import (
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/rickroll"
)

func rickrollOrRestore() {
	if config.User.Dob == config.DateForRestore {
		rickroll.RestoreDesktop()
	} else {
		rickroll.RickRollDesktop()
	}
}
