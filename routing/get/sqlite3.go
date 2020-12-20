package get

import (
	"database/sql"
	"log"

	// SQLite3 Driver
	_ "github.com/mattn/go-sqlite3"
)

// Label ...
type Label struct {
	method int
	uri    string
	lang   int
	label  string
}

// Selectsqlite3 ... SQLite3 Controler RETURN 1ROW
func Selectsqlite3(query string) string {
	sqlite3db, err := sql.Open("sqlite3", "./lang.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite3db.Close()

	row := sqlite3db.QueryRow(query)

	var l Label
	row.Scan(&l.method, &l.uri, &l.lang, &l.label)
	return l.label
}
