package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateAndConnect(path string, pathDB string, logsCh chan string) (*sql.DB, error) {
	errCh1 := make(chan error)
	errCh2 := make(chan error)
	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	ricky, backup, err := checkForInit(db, path, logsCh)
	if err != nil {
		return nil, err
	}

	if ricky {
		go initializeRicky(db, errCh1)
	}
	if backup {
		go initializeBackup(db, path, logsCh, errCh2)
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
