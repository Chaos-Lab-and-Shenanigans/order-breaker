package sqlite3

import (
	"fmt"
	"os"

	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
)

// Checks if initialzation is required for both tables depending on lenght in DB vs Folder
func checkForInit() (bool, bool, error) {
	var backup bool
	var ricky bool
	var lenR int //Total length including the first initializer
	var lenB int //Same
	db := config.Cfg.DB
	path := config.PATH

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

	//lenB-1 bcz first element in db is the intializer
	config.Cfg.LogsCh <- fmt.Sprintf("Items in db(exlcuding program files):     %v", lenB-1)
	config.Cfg.LogsCh <- fmt.Sprintf("Items in folder(including program files): %v", len(items))

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

	if lenR != len(config.LYRICS) {
		ricky = true
	}
	return ricky, backup, nil
}
