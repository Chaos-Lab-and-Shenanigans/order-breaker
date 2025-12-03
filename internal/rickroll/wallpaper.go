package rickroll

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
	"github.com/reujab/wallpaper"
)

func restoreWallpaper(errCh chan error) {
	if config.BackupWall == "" {
		errCh <- nil
		return
	}

	if strings.Contains(strings.ToLower(config.BackupWall), "transcodedwallpaper") {
		errCh <- fmt.Errorf("Change the wallpaper back yourself\n")
		return
	}

	err := wallpaper.SetFromFile(config.BackupWall)
	if err != nil {
		errCh <- fmt.Errorf("Error restoring wallpaper: %v", err)
		return
	}

	errCh <- nil
}

func setWallpaper(errCh chan error) {
	rickyWallPath := filepath.Join(config.Cfg.Path, "wall.png")
	err := backupWallpaper(rickyWallPath)
	if errors.Is(err, config.RetErr) {
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

	if currentWallPath == rickyWallPath {
		return config.RetErr
	}

	config.BackupWall = currentWallPath
	return nil
}
