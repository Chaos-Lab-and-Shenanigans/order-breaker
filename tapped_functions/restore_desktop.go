package tappedfunctions

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func TapRestoreDesktop(db *sql.DB, x chan string, path string, pathDB string) func() {
	return func() {
		items, err := os.ReadDir(path)
		if err != nil {
			x <- fmt.Sprintf("Error occured while listing \"%v\": %v", path, err)
			return
		}

		for _, item := range items {
			name := item.Name()
			if name == "backupob.db" {
				continue
			}
			id, err := getID(name)
			if err != nil {
				x <- fmt.Sprintf("Error getting ID from file: %v\nGet rickrolled first mf", err)
				return
			}

			//fmt.Printf("Counter: %v\n", i)
			var ogName string
			cmd := fmt.Sprintf("SELECT body FROM backup WHERE id = %v", id)
			row := db.QueryRow(cmd)
			err = row.Scan(&ogName)
			if err != nil {
				x <- fmt.Sprintf("Error while scanning row#%v: %v", id, err)
				return
			}

			oldName := filepath.Join(path, name)
			ogName = filepath.Join(path, ogName)

			err = os.Rename(oldName, ogName)
			if err != nil {
				x <- fmt.Sprintf("Error while renaming file: %v", err)
				return
			}
		}
		x <- "Successfully recovered the madness"
	}
}

func getID(name string) (int, error) {
	idS, _, found := strings.Cut(name, "#")
	if !found {
		return 0, fmt.Errorf("The file \"%v\" doesn't contain ID", name)
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, fmt.Errorf("The file \"%v\" doesn't contain ID", name)
	}

	return id, nil
}
