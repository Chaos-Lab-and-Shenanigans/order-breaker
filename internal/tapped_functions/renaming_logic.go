package tappedfunctions

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func rickrollNames(db *sql.DB, path string, errCh chan error) {
	items, err := os.ReadDir(path)
	if err != nil {
		errCh <- fmt.Errorf("Error occured while reading directory: %v", err)
		return
	}

	i := 0
	for _, item := range items {
		name := item.Name()
		if name == "backupob.db" || name == APP_NAME { //Skipping program files
			continue
		}

		i += 1
		if AlreadyMessedUp(name) {
			fmt.Printf("File \"%v\" already messed up\nSkipping...\n", name)
			continue
		}

		var rickName string
		id := (i-1)%limit + 1 //Skipping ID no 0 while looping over the limit
		cmd := fmt.Sprintf("SELECT body FROM ricky WHERE id = %v", id)
		row := db.QueryRow(cmd)
		err = row.Scan(&rickName)
		if err != nil {
			errCh <- fmt.Errorf("Error occured while scanning row#%v: %v", i, err)
			return
		}

		rickName = strconv.Itoa(i) + sep + rickName //For sorting correctly(required for restoring back correctly)
		ogName := filepath.Join(path, name)
		newName := filepath.Join(path, rickName)
		err = os.Rename(ogName, newName)
		if err != nil {
			errCh <- fmt.Errorf("Error renaming file: %v", err)
			return
		}
	}
	errCh <- nil
}

func AlreadyMessedUp(name string) bool {
	return strings.Contains(name, sep)
}

func restoreNames(db *sql.DB, path string, errCh chan error) {
	items, err := os.ReadDir(path)
	if err != nil {
		errCh <- fmt.Errorf("Error occured while listing \"%v\": %v", path, err)
		return
	}

	for _, item := range items {
		name := item.Name()
		if name == "backupob.db" || name == APP_NAME { //Skipping program files
			continue
		}

		id, err := getID(name)
		if err != nil { //If id is not in name, then its normal file not renamed one so skip it
			continue
		}

		var ogName string
		cmd := fmt.Sprintf("SELECT body FROM backup WHERE id = %v", id)
		row := db.QueryRow(cmd)
		err = row.Scan(&ogName)
		if err != nil {
			errCh <- fmt.Errorf("Error while scanning row#%v: %v", id, err)
			return
		}

		oldName := filepath.Join(path, name)
		ogName = filepath.Join(path, ogName)

		err = os.Rename(oldName, ogName)
		if err != nil {
			errCh <- fmt.Errorf("Error while renaming file: %v", err)
			return
		}
	}
	errCh <- nil
}

func getID(name string) (int, error) {
	idS, _, found := strings.Cut(name, sep)
	if !found {
		return 0, fmt.Errorf("The file \"%v\" doesn't contain ID", name)
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, fmt.Errorf("The file \"%v\" doesn't contain ID", name)
	}

	return id, nil
}
