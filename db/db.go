package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	// importing the package solely for its side effects —
	// not to use any of its exported identifiers (functions, types, etc.) directly in code
)

var sqlDatabase *sql.DB

func InitDB() {
	var err error
	sqlDatabase, err = sql.Open("sqlite3", "api.db")
	//sql.Open() => driverName must be "sqlite3" to make use of package
	//sql.Open() => dataSourceName must must end in .db

	if err != nil {
		panic("Could not connect to the database")
	}

	defer sqlDatabase.Close()

	sqlDatabase.SetMaxOpenConns(10)
	sqlDatabase.SetMaxIdleConns(5)

}
