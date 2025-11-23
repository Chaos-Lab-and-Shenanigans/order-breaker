package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateAndConnect(x chan string, path string, pathDB string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	ricky, backup, err := checkForInit(db, path)
	if err != nil {
		return nil, err
	}

	if ricky {
		err = initializeRicky(db)
		if err != nil {
			return nil, err
		}
	}
	if backup {
		err = initializeBackup(db, path, x)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
