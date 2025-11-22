package sqlite3

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	limit  = 10
	lyrics = []string{
		"NEVER GONNA", "GIVE YOU UP",
		"NEVER GONNA", "LET YOU DOWN",
		"NEVER GONNA", "RUN AROUND", "AND DESERT YOU",
		"NEVER GONNA", "MAKE YOU CRY",
		"NEVER GONNA", "SAY GOODBYE",
		"NEVER GONNA", "TELL A LIE", "AND HURT YOU",
	}
)

func CreateAndConnect(pathDB string, x chan string) (*sql.DB, error) {
	path := getParentDir(pathDB) //get directory's path
	_, err := os.OpenFile(pathDB, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

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

func getParentDir(path string) string {
	paths := strings.Split(path, "\\")
	pathWithoutChildArray := paths[0 : len(paths)-1] //Without last element

	pathWithoutChild := strings.Join(pathWithoutChildArray, "\\")
	return pathWithoutChild
}
