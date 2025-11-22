package tappedfunctions

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func TapRickRollDesktop(db *sql.DB, x chan string, path string) func() {
	return func() {
		items, err := os.ReadDir(path)
		if err != nil {
			x <- fmt.Sprintf("Error occured while reading directory: %v", err)
			return
		}

		i := 0
		for _, item := range items {
			name := item.Name()
			if name == "backupob.db" { //Skipping db file
				continue
			}

			i += 1
			var rickName string
			cmd := fmt.Sprintf("SELECT body FROM ricky WHERE id = %v", i)
			row := db.QueryRow(cmd)
			err = row.Scan(&rickName)
			if err != nil {
				x <- fmt.Sprintf("Error occured while scanning row#%v: %v", i, err)
				return
			}

			rickName = strconv.Itoa(i) + "#" + rickName //For sorting correctly
			ogName := filepath.Join(path, name)
			newName := filepath.Join(path, rickName)
			err = os.Rename(ogName, newName)
			if err != nil {
				x <- fmt.Sprintf("Error renaming file: %v", err)
				return
			}
		}

		x <- "Check out your desktop brother 🙂"
	}
}
