package tappedfunctions

import (
	"fmt"

	"github.com/reujab/wallpaper"
)

const (
	limit = 14
	sep   = "~"
)

var (
	APP_NAME           = "prank.exe"
	backupWall         = ""
	currentWallPath, _ = wallpaper.Get()
	retErr             = fmt.Errorf("return")
)
