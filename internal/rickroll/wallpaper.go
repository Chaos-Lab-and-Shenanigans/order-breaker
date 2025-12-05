package rickroll

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
	"github.com/reujab/wallpaper"
)

func restoreWallpaper(errCh chan error) {
	if config.BackupWall == "" {
		errCh <- nil
		return
	}

	err := wallpaper.SetFromFile(config.BackupWall)
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}

func setWallpaper(errCh chan error) {
	rickyWallPath := filepath.Join(config.PATH, "wall.png")
	err := backupWallpaper(rickyWallPath)
	if errors.Is(err, config.WallAlreadyFkedErr) {
		errCh <- nil
		return
	}
	if err != nil {
		errCh <- err
		return
	}

	config.Cfg.LogsCh <- fmt.Sprintf("Backup wallpaper: %v", config.BackupWall)

	err = os.WriteFile(rickyWallPath, *config.Cfg.RickyWall, 0666)
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

	if strings.Contains(strings.ToLower(currentWallPath), "transcodedwallpaper") {
		config.IsSlideshowWall = true
		return nil
	}

	if currentWallPath == rickyWallPath {
		return config.WallAlreadyFkedErr
	}

	config.BackupWall = currentWallPath
	fmt.Printf("Wall: %v\n", config.BackupWall)

	return nil
}
