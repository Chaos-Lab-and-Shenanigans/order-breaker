package sqlite3

import (
	"database/sql"
	"fmt"
	"os"
)

// Checks if initialzation is required for both tables depending on lenght in DB vs Folder
func checkForInit(db *sql.DB, path string, logsCh chan string) (bool, bool, error) {
	var backup bool
	var ricky bool
	var lenR int
	var lenB int

	//Backup table's check
	row := db.QueryRow("SELECT COUNT(*) FROM backup")
	err := row.Scan(&lenB)
	if err != nil { //Means table doesn't exist, yet
		backup = true
	}

	items, err := os.ReadDir(path)
	if err != nil {
		err = fmt.Errorf("Error while listing \"%v\" contents: %v", path, err)
		return ricky, backup, err
	}

	logsCh <- fmt.Sprintf("\nLength in db: %v\nLength in folder: %v\n", lenB, len(items))

	if lenB != len(items) {
		backup = true
	}

	//Ricky table's check
	row = db.QueryRow("SELECT COUNT(*) FROM ricky")
	err = row.Scan(&lenR)
	if err != nil { //Means table doesn't exist, yet
		ricky = true
		return ricky, backup, nil
	}

	if lenR != len(lyrics) {
		ricky = true
	}
	return ricky, backup, nil
}
