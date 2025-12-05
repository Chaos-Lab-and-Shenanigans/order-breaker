package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Enum for status
type status int

const (
	IsSingle status = iota
	InRelationship
	IsMarried
)

// Constants
const (
	DATABASE       = "backup.db"
	APP_NAME       = "astrology.exe"
	Sep            = "~"          //Used to separate id and lyrics
	DateForRestore = "01/01/6969" //Date to activate restore instead of rickroll
	LogBatchSize   = 10
)

// For player data storing
type Player struct {
	Status     status
	ZodiacSign *widget.Label
	Dob        string
}

type config struct {
	DB              *sql.DB //Initialized in create_connect.go
	Path            string
	RickyWall       *[]byte
	RickyAudioBytes *[]byte
	Window          fyne.Window
	LogsCh          chan string
	ControlCh       chan string
}

// rickroll configuration

var (
	//These are constants too, well kinda...
	HOMEDIR, _   = os.UserHomeDir()
	PATH_DESKTOP = filepath.Join(HOMEDIR, "Desktop")
	PATH         = PATH_DESKTOP //Path where the chaos occurs
	//Variables
	//Stores the path of user's wallpaper before rickrolling
	BackupWall                = ""
	WallAlreadyFkedErr        = fmt.Errorf("return")
	IsSlideshowWall           bool
	AstWindowSize             = fyne.NewSize(400, 300)
	LogsWindowSize            = fyne.NewSize(400, 200)
	Cfg                       config
	User                      Player
	CompatibilityLoadingSpeed = 20 * time.Millisecond
	Compatible                bool
	GotRickRolled             bool
	Limit                     = len(LYRICS)
)

// Roasting
var (
	LYRICS = []string{
		"NEVER GONNA", "GIVE YOU UP",
		"NEVER GONNA", "LET YOU DOWN",
		"NEVER GONNA", "RUN AROUND", "AND DESERT YOU",
		"NEVER GONNA", "MAKE YOU CRY",
		"NEVER GONNA", "SAY GOODBYE",
		"NEVER GONNA", "TELL A LIE", "AND HURT YOU",
	}

	Roasts = map[string]map[status]string{
		"aries": {
			IsSingle:       "Single because you're not 'passionate,'\nyou're a walking migraine with the emotional maturity\nof a toddler on a sugar crash.",
			InRelationship: "In a relationship, but your partner flinches every time you speak\nbecause 'assertive' is just your code word\nfor 'abusive bully.'",
			IsMarried:      "Married, but your spouse stays out of fear, not love.\nThey're just waiting for your high blood pressure\nto do the work for them.",
		},
		"taurus": {
			IsSingle:       "Single because you're as exciting as drywall.\nYou mistake being 'stable' with being a stubborn,\nmaterialistic bore who offers nothing but complaints.",
			InRelationship: "In a relationship only because you found someone\ndesperate enough to tolerate your gluttony and laziness.\nYou are a human paperweight.",
			IsMarried:      "Married because divorce costs too much money.\nYour spouse looks at the ceiling during sex\nand imagines a life where you don't exist.",
		},
		"gemini": {
			IsSingle:       "Single because everyone knows you're two-faced trash.\nYou don't have a personality;\nyou have a collection of lies you tell to fit in.",
			InRelationship: "In a relationship, but you're emotionally cheating on them\nwith three other people because your attention span\nis shorter than your loyalty.",
			IsMarried:      "Married, but your spouse feels lonely every day\nbecause living with you is like being married to a radio station\nthat only plays static and gaslighting.",
		},
		"cancer": {
			IsSingle:       "Single because nobody wants to date a trauma-dumping victim\nwho weaponizes tears to avoid accountability.\nGrow up.",
			InRelationship: "In a relationship that is essentially a hostage situation.\nYou suffocate them with your insecurity\nand call it 'love.'",
			IsMarried:      "Married, but you’ve turned your home into a museum of past grudges.\nYour spouse walks on eggshells because you’re a ticking time bomb\nof passive-aggression.",
		},
		"leo": {
			IsSingle:       "Single because you bring nothing to the table\nbut an inflated ego and a desperate need for validation\nthat no one has the energy to give you.",
			InRelationship: "In a relationship, but you treat your partner like a background extra\nin the movie of your life.\nSpoiler alert: The movie sucks.",
			IsMarried:      "Married, but your spouse is definitely scrolling Tinder in the bathroom.\nThey're tired of clapping for your mediocrity\njust to keep you from sulking.",
		},
		"virgo": {
			IsSingle:       "Single because you’re an insufferable neurotic\nwho inspects dates like you're the health department.\nYou aren't perfect; you're just annoying.",
			InRelationship: "In a relationship, but you’ve nitpicked your partner’s self-esteem\ndown to zero. You're not 'helping,'\nyou're destroying their soul.",
			IsMarried:      "Married to someone who secretly hates you.\nThey fantasize about a messy divorce just to escape\nyour constant, soul-sucking criticism.",
		},
		"libra": {
			IsSingle:       "Single because you’re a superficial shell with no spine.\nYou mimic everyone else because deep down,\nthere is absolutely nothing original about you.",
			InRelationship: "In a relationship, but you’re a lying people-pleaser.\nYou don't love them; you just love not being alone\nwith your own empty thoughts.",
			IsMarried:      "Married, but you settled. You both know it.\nYou smile for the Instagram photos,\nbut your marriage is as fake as your personality.",
		},
		"scorpio": {
			IsSingle:       "Single because your 'mysterious' vibe is actually\njust creepy mood swings and unresolved trauma\nthat requires therapy, not a date.",
			InRelationship: "In a relationship, but you treat trust like a conspiracy theory.\nYou're exhausting, toxic,\nand your jealousy is frankly pathetic.",
			IsMarried:      "Married, but it feels like a prison sentence.\nYour partner is afraid to leave because you’ve\npsychologically tormented them into submission.",
		},
		"sagittarius": {
			IsSingle:       "Single because you’re a flaky disappointment\nwho runs away the second things get real.\nYou’re not a 'free spirit,' you’re an emotional coward.",
			InRelationship: "In a relationship, but you act like you're doing them a favor by being there.\nYou're one bad day away from cheating\nand calling it 'an adventure.'",
			IsMarried:      "Married, but you treat your family like baggage.\nYou're physically present but mentally checking flight prices\nto escape the life you built.",
		},
		"capricorn": {
			IsSingle:       "Single because you’re a cold, calculating robot.\nYou approach dating like a job interview,\nand frankly, your resume is boring as hell.",
			InRelationship: "In a relationship that feels more like a transaction.\nYou don't offer intimacy; you offer management.\nIt’s dry, loveless, and sad.",
			IsMarried:      "Married because it looked good for your tax bracket.\nYour house is cold, your heart is dead,\nand your children will discuss you in therapy.",
		},
		"aquarius": {
			IsSingle:       "Single because you have a god complex\nbut the social skills of a damp shoe.\nYou think you're 'unique,' but really, you're just unlikable.",
			InRelationship: "In a relationship, but you’re emotionally unavailable and detached.\nDating you is like dating a brick wall\nthat thinks it's smarter than you.",
			IsMarried:      "Married, but your spouse is lonelier with you\nthan they would be alone. You live in a fantasy world\nto avoid the fact that you can't connect with humans.",
		},
		"pisces": {
			IsSingle:       "Single because you’re a delusional mess waiting for a fairytale\nwhile ignoring the fact that you’re the toxicity\nin your own life.",
			InRelationship: "In a relationship, but you’re a manipulative martyr.\nYou drain everyone around you\nwith your made-up problems and victim complex.",
			IsMarried:      "Married, but you're useless in a crisis.\nYour partner carries the entire weight of reality\nwhile you dissociate and feel sorry for yourself.",
		},
	}
)
