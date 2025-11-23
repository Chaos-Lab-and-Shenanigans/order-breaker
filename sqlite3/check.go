package sqlite3

import (
	"database/sql"
	"fmt"
	"os"
)

// Checks if initialzation is required for both tables
func checkForInit(db *sql.DB, path string) (bool, bool, error) {
	var backup bool
	var ricky bool
	var lenR int
	var lenB int
	//Backup table's check
	row := db.QueryRow("SELECT COUNT(*) FROM backup")
	err := row.Scan(&lenB)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		backup = true
	}

	items, err := os.ReadDir(path)
	if err != nil {
		err = fmt.Errorf("Error while listing \"%v\" contents: %v", path, err)
		return ricky, backup, err
	}

	//length-1 bcz the first row is initializer and does not contain data
	fmt.Printf("\nLength in db: %v\nLength in folder: %v\n", lenB, len(items))

	if lenB != len(items) {
		backup = true
	}

	//Ricky table's check
	row = db.QueryRow("SELECT COUNT(*) FROM ricky")
	err = row.Scan(&lenR)
	if err != nil {
		err = fmt.Errorf("Error executing query: %v", err)
		ricky = true
		return ricky, backup, nil
	}

	if lenR != len(lyrics) {
		ricky = true
	}
	return ricky, backup, nil
}
