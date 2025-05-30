package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	// importing the package solely for its side effects —
	// not to use any of its exported identifiers (functions, types, etc.) directly in code
)

var SqlDatabase *sql.DB

func InitDB() {
	var err error
	SqlDatabase, err = sql.Open("sqlite3", "api.db")
	//sql.Open() => driverName must be "sqlite3" to make use of package
	//sql.Open() => dataSourceName must must end in .db

	if err != nil {
		panic("Could not connect to the database")
	}

	// defer SqlDatabase.Close()

	SqlDatabase.SetMaxOpenConns(10)
	SqlDatabase.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err := SqlDatabase.Exec(createUsersTable)

	if err != nil {
		panic("Unable to create users table")
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	// statement, err := sqlDatabase.Prepare(createEventsTable)

	// if err != nil {
	// 	panic("Unable to create events table")
	// }

	_, err = SqlDatabase.Exec(createEventsTable)
	if err != nil {
		panic("Unable to create events table")
	}

	createRegistrationTable := `CREATE TABLE IF NOT EXISTS registration(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = SqlDatabase.Exec(createRegistrationTable)
	if err != nil {
		panic("Unable to create registration table")
	}
}
