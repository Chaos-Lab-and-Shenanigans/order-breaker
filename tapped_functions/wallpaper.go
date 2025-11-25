package tappedfunctions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/reujab/wallpaper"
)

func restoreWallpaper(errCh chan error) {
	fmt.Printf("Restore wallpaper: %v\n", backupWall)
	if backupWall == "" {
		errCh <- nil
		return
	}

	err := wallpaper.SetFromFile(backupWall)
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}

func setWallpaper(rickyWall []byte, path string, errCh chan error) {
	fullName := filepath.Join(path, "wall.png")
	fmt.Printf("Wallpaper: %v", fullName)
	err := backupWallpaper(fullName)
	if err != nil {
		errCh <- err
		return
	}

	err = os.WriteFile(fullName, rickyWall, 0666)
	if err != nil {
		errCh <- err
		return
	}

	err = wallpaper.SetFromFile(fullName)
	if err != nil {
		errCh <- err
		return
	}

	err = os.Remove(fullName)
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}

func backupWallpaper(rickyWall string) error {
	currentWall, err := wallpaper.Get()
	if err != nil {
		return err
	}

	if currentWall == rickyWall {
		return nil
	}

	backupWall = currentWall
	return nil
}
