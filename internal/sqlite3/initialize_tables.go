package sqlite3

import (
	"fmt"
	"os"

	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/rickroll"
)

// Create and initlalize ricky table
func initializeRicky(errCh chan error) {
	db := config.Cfg.DB
	db.Exec("DROP TABLE ricky")
	_, err := db.Exec("CREATE TABLE ricky(id INTEGER PRIMARY KEY, body TEXT NOT NULL)")
	if err != nil {
		errCh <- err
		return
	}

	_, err = db.Exec("INSERT INTO ricky(id, body) VALUES (0, \"Start of auto increment\")")
	if err != nil {
		errCh <- err
		return
	}

	for _, words := range config.LYRICS {
		query := fmt.Sprintf("INSERT INTO ricky(body) VALUES (\"%v\")", words)
		_, err = db.Exec(query)
		if err != nil {
			errCh <- err
			return
		}
	}
	errCh <- nil
}

func initializeBackup(errCh chan error) {
	items, err := os.ReadDir(config.PATH)
	if err != nil {
		errCh <- err
	}

	index := getNonMessedIndex(items)

	if index <= 1 {
		config.Cfg.LogsCh <- "Recreating backup"
		err = recreateBackup(items)
		errCh <- err
	} else {
		config.Cfg.LogsCh <- "Adding to existing backup"
		err = addToExistingTable(items, index)
		errCh <- err
	}
	errCh <- nil
}

func addToExistingTable(items []os.DirEntry, index int) error {
	neededItems := items[index-1:] //neededItems contains only those items that need to be stored
	for i, item := range neededItems {
		name := item.Name()
		if name == config.DATABASE || name == config.APP_NAME {
			continue
		}

		query := fmt.Sprintf("INSERT INTO backup(body) VALUES (\"%v\")", name)
		_, err := config.Cfg.DB.Exec(query)
		if err != nil {
			return err
		}
		config.Cfg.LogsCh <- fmt.Sprintf("Item %v: %v added to backup\n", i+1, name)
	}
	return nil
}

func getNonMessedIndex(items []os.DirEntry) int { //returns the index of item not messed up
	indexForNonMessed := 0
	for _, item := range items {
		name := item.Name()

		if name == config.APP_NAME || name == config.DATABASE { //skip without incrementing if app's file
			continue
		}

		indexForNonMessed += 1
		if rickroll.AlreadyMessedUp(item.Name()) {
			continue
		}
		return indexForNonMessed
	}
	return indexForNonMessed
}

// Create and initialize backup table
func recreateBackup(items []os.DirEntry) error {
	db := config.Cfg.DB
	db.Exec("DROP TABLE backup")
	_, err := db.Exec("CREATE TABLE backup(id INTEGER PRIMARY KEY, body TEXT NOT NULL)")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO backup(id, body) VALUES (0, \"Initialized for auto increment\")")
	if err != nil {
		return err
	}

	//Read and create backup in database
	err = addToExistingTable(items, 1)
	if err != nil {
		return err
	}

	config.Cfg.LogsCh <- "Backup created successfully"
	return nil
}
