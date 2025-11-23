package tappedfunctions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/reujab/wallpaper"
)

func restoreWallpaper() error {
	fmt.Printf("Restore wallpaper: %v\n", backupWall)
	if backupWall == "" {
		return nil
	}
	err := wallpaper.SetFromFile(backupWall)
	if err != nil {
		return err
	}
	return nil
}

func setWallpaper(rickyWall []byte, path string) error {
	fullName := filepath.Join(path, "wall.png")
	fmt.Printf("Wallpaper: %v", fullName)
	err := backupWallpaper(fullName)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullName, rickyWall, 0666)
	if err != nil {
		return err
	}

	err = wallpaper.SetFromFile(fullName)
	if err != nil {
		return err
	}

	err = os.Remove(fullName)
	if err != nil {
		return err
	}

	return nil
}

func backupWallpaper(rickyWall string) error {
	ogWall, err := wallpaper.Get()
	if err != nil {
		return err
	}

	if rickyWall == ogWall {
		return nil
	}

	backupWall = ogWall
	return nil
}
