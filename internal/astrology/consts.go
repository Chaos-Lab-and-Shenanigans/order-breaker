package handleastrology

import "fyne.io/fyne/v2/widget"

const (
	dateToCheckForRecovery = "05/15/2007"
)

type Player struct {
	single       bool
	married      bool
	relationship bool
	zodiacSign   *widget.Label
}

var player Player
