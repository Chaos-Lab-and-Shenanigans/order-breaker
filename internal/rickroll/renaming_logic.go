package rickroll

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
)

func rickrollNames(errCh chan error) {
	dbName := config.DATABASE
	items, err := os.ReadDir(config.PATH)
	if err != nil {
		errCh <- fmt.Errorf("Error occured while reading directory: %v", err)
		return
	}

	i := 0
	for _, item := range items {
		name := item.Name()
		if name == dbName || name == config.APP_NAME { //Skipping program files
			continue
		}

		i += 1
		if AlreadyMessedUp(name) {
			config.Cfg.LogsCh <- fmt.Sprintf("File \"%v\" already messed up\nSkipping...\n", name)
			continue
		}

		var rickName string
		id := (i-1)%config.Limit + 1 //Skipping ID no 0 while looping over the limit
		cmd := fmt.Sprintf("SELECT body FROM ricky WHERE id = %v", id)
		row := config.Cfg.DB.QueryRow(cmd)
		err = row.Scan(&rickName)
		if err != nil {
			errCh <- fmt.Errorf("Error occured while scanning row#%v: %v", i, err)
			return
		}

		rickName = strconv.Itoa(i) + config.Sep + rickName //For sorting correctly(required for restoring back correctly)
		ogName := filepath.Join(config.PATH, name)
		newName := filepath.Join(config.PATH, rickName)
		err = os.Rename(ogName, newName)
		if err != nil {
			errCh <- fmt.Errorf("Error renaming file: %v", err)
			return
		}
	}
	errCh <- nil
}

func AlreadyMessedUp(name string) bool {
	return strings.Contains(name, config.Sep)
}

func restoreNames(errCh chan error) {
	dbName := config.DATABASE
	items, err := os.ReadDir(config.PATH)
	if err != nil {
		errCh <- fmt.Errorf("Error occured while listing \"%v\": %v", config.PATH, err)
		return
	}

	for _, item := range items {
		name := item.Name()
		if name == dbName || name == config.APP_NAME { //Skipping program files
			continue
		}

		id, err := getID(name)
		if err != nil { //If id is not in name, then its normal file not renamed one so skip it
			continue
		}

		var realName string
		cmd := fmt.Sprintf("SELECT body FROM backup WHERE id = %v", id)
		row := config.Cfg.DB.QueryRow(cmd)
		err = row.Scan(&realName)
		if err != nil {
			errCh <- fmt.Errorf("Error while scanning row#%v: %v", id, err)
			return
		}

		oldName := filepath.Join(config.PATH, name)
		realName = filepath.Join(config.PATH, realName)

		err = os.Rename(oldName, realName)
		if err != nil {
			errCh <- fmt.Errorf("Error while renaming file: %v", err)
			return
		}
	}
	errCh <- nil
}

func getID(name string) (int, error) {
	idS, _, found := strings.Cut(name, config.Sep)
	if !found {
		return 0, fmt.Errorf("The file \"%v\" doesn't contain ID", name)
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, fmt.Errorf("The file \"%v\" doesn't contain ID", name)
	}

	return id, nil
}
