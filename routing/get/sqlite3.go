package get

import (
	"database/sql"
	"log"

	// SQLite3 Driver
	_ "github.com/mattn/go-sqlite3"
)

// Label ... getting labelData
type Label struct {
	method int
	uri    string
	lang   int
	label  string
}

// Datalabel ... SQLite3 Controler RETURN 1ROW
func Datalabel(uri string, lang int) string {

	query := "SELECT label FROM mst_label WHERE method = 1 AND uri = ? AND mst_label.lang = ?;"
	sqlite3db, err := sql.Open("sqlite3", "./lang.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite3db.Close()

	row := sqlite3db.QueryRow(query, uri, lang)

	var l Label
	row.Scan(&l.label)
	return l.label
}
