package sqlite3

import (
	"database/sql"
	"fmt"
	"os"
)

// Create and initlalize ricky table
func initializeRicky(db *sql.DB) error {

	db.Exec("DROP TABLE ricky")
	_, err := db.Exec("CREATE TABLE ricky(id INTEGER PRIMARY KEY, body TEXT NOT NULL)")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO ricky(id, body) VALUES (0, \"Start of auto increment\")")
	if err != nil {
		return err
	}

	for _, name := range lyrics {
		cmd := fmt.Sprintf("INSERT INTO ricky(body) VALUES (\"%v\")", name)
		_, err = db.Exec(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create and initialize backup table
func initializeBackup(db *sql.DB, path string, x chan string) error {
	db.Exec("DROP TABLE backup")
	_, err := db.Exec("CREATE TABLE backup(id INTEGER PRIMARY KEY, body TEXT NOT NULL)")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO backup(id, body) VALUES (0, \"Initialized for auto increment\")")
	if err != nil {
		return err
	}

	//Read and create backup in db
	items, err := os.ReadDir(path)
	if err != nil {
		err = fmt.Errorf("Error occured while reading directory: %v", err)
		return err
	}

	i := 0
	for _, item := range items {
		name := item.Name()
		if name == "backupob.db" {
			continue
		}

		i += 1
		cmd := fmt.Sprintf("INSERT INTO backup(body) VALUES (\"%v\")", name)
		_, err = db.Exec(cmd)
		if err != nil {
			err = fmt.Errorf("Error occured while creating backup: %v", err)
			return err
		}
	}

	x <- "Backup created successfully"
	return nil
}
