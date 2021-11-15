package connection

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./store.db3")
	HandleErr(err)
	createTableQuery := `create table if not exists students(
		id integer not null primary key autoincrement,
		name text,
		age integer
	)`

	res, err := Db.Exec(createTableQuery)
	HandleErr(err)
	fmt.Println(res)
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
