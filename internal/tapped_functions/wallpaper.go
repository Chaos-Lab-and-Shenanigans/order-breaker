package tappedfunctions

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/reujab/wallpaper"
)

func restoreWallpaper(errCh chan error) {
	if backupWall == "" {
		errCh <- nil
		return
	}

	if strings.Contains(strings.ToLower(backupWall), "transcodedwallpaper") {
		errCh <- fmt.Errorf("Change the wallpaper back yourself\n")
		return
	}

	err := wallpaper.SetFromFile(backupWall)
	if err != nil {
		errCh <- fmt.Errorf("Error restoring wallpaper: %v", err)
		return
	}

	errCh <- nil
}

func setWallpaper(rickyWall []byte, path string, logsCh chan string, errCh chan error) {
	rickyWallPath := filepath.Join(path, "wall.png")
	logsCh <- fmt.Sprintf("Wallpaper: %v\n", rickyWallPath)
	err := backupWallpaper(rickyWallPath)
	if errors.Is(err, retErr) {
		errCh <- nil
		return
	}
	if err != nil {
		errCh <- err
		return
	}

	err = os.WriteFile(rickyWallPath, rickyWall, 0666)
	if err != nil {
		errCh <- err
		return
	}

	err = wallpaper.SetFromFile(rickyWallPath)
	if err != nil {
		errCh <- err
		return
	}

	err = os.Remove(rickyWallPath)
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}

func backupWallpaper(rickyWallPath string) error {
	currentWallPath, err := wallpaper.Get()
	if err != nil {
		return err
	}

	if currentWallPath == rickyWallPath {
		return retErr
	}

	backupWall = currentWallPath
	return nil
}
