package sqlite3

import (
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/Chaos-Lab-and-Shenanigans/astrology/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func CreateAndConnect() (*sql.DB, error) {
	pathDB := filepath.Join(config.PATH, config.DATABASE)
	errCh1 := make(chan error)
	errCh2 := make(chan error)

	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		config.Cfg.LogsCh <- fmt.Sprintf("Database path: %v", pathDB)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	config.Cfg.DB = db

	ricky, backup, err := checkForInit()
	if err != nil {
		return nil, err
	}

	if ricky {
		go initializeRicky(errCh1)
	}
	if backup {
		go initializeBackup(errCh2)
	}

	//return errors, if any
	if ricky {
		err := <-errCh1
		if err != nil {
			return nil, err
		}
	}

	if backup {
		err := <-errCh2
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
