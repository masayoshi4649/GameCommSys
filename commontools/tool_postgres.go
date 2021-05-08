package commontools

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PsqlSelect() {
	db, err := sql.Open("postgres", "host="+Fixedparam("postgreshost")+"port=5432 user=postgres password=postgres dbname=gamecommsysdb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
}
