package commontools

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

// FixedData ... SQLiteマスタ定数固定値
type FixedData struct {
	param string
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

// Fixedparam ... 固定値をDBより取得する
func Fixedparam(title string) string {
	query := "SELECT param FROM mst_fixeddata WHERE title = ?;"
	sqlite3db, err := sql.Open("sqlite3", "./lang.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite3db.Close()

	row := sqlite3db.QueryRow(query, title)

	var p FixedData
	row.Scan(&p.param)
	return p.param
}
