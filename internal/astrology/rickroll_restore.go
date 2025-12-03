package astrology

import (
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/rickroll"
)

func rickrollOrRestore() {
	if config.User.Dob == config.DateForRecovery {
		rickroll.RestoreDesktop()
	} else {
		rickroll.RickRollDesktop()
	}
}
