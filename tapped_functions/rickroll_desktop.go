package tappedfunctions

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func TapRickRollDesktop(db *sql.DB, x chan string, path string, rickyWall []byte) func() {
	return func() {
		errCh1 := make(chan error)
		errCh2 := make(chan error)

		go setWallpaper(rickyWall, path, errCh1)
		go rickrollNames(db, path, errCh2)

		err := <-errCh1
		if err != nil {
			x <- fmt.Sprintf("Error occured while setting wallpaper: %v", err)
			return
		}

		err = <-errCh2
		if err != nil {
			x <- fmt.Sprintf("Error occured while rickrolling names: %v", err)
			return
		}

		x <- "Check out your desktop brother 🙂"
	}
}

func rickrollNames(db *sql.DB, path string, errCh chan error) {
	items, err := os.ReadDir(path)
	if err != nil {
		errCh <- fmt.Errorf("Error occured while reading directory: %v", err)
		return
	}

	i := 0
	for _, item := range items {
		name := item.Name()
		if name == "backupob.db" { //Skipping program files
			continue
		}

		if alreadyMessedUp(name) {
			fmt.Printf("File \"%v\" already messed up\nExiting...\n", name)
			return
		}

		i += 1
		var rickName string
		id := (i-1)%limit + 1 //Skipping ID no 0
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

func alreadyMessedUp(name string) bool {
	return strings.Contains(name, sep)
}
