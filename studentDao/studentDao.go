package studentDao

import (
	"database/sql"
	"fmt"

	conn "github.com/elmubarak/go_crud/connection"
)

func FetchOneFromDb(id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from students where id='%v'", id)
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func FetchAllFromDb() (*sql.Rows, error) {
	query := "select * from students"
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func UpdateOneFromDb(name string, age int, id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("update students set name='%v', age='%v'  where id='%v'", name, age, id)
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func DeleteOneFromDb(id int64) (*sql.Rows, error) {
	query := fmt.Sprintf("delete from students where id='%v'", id)
	rows, err := conn.Db.Query(query)
	HandleErr(err)
	return rows, err
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
