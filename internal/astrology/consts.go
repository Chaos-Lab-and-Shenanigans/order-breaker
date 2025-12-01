package astrology

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

const (
	dateForRecovery = "01/01/6969"
)

type Player struct {
	isSingle       bool
	isMarried      bool
	inRelationship bool
	zodiacSign     *widget.Label
	dob            string
}

type config struct {
	DB              *sql.DB
	Path            string
	PathDB          string
	RickyWall       *[]byte
	RickyAudioBytes *[]byte
	Window          fyne.Window
	LogsCh          chan string
	RestartCh       chan string
}

var cfg config
var player Player
