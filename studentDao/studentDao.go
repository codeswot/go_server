package studentDao

import (
	"database/sql"
	"fmt"

	conn "github.com/elmubarak/go_server/connection"
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

func UpdateOneFromDb(id int64, fName string, lName string, email string, gender string, age int) (*sql.Rows, error) {
	query := fmt.Sprintf("update students set fName='%v',lName='%v', gender = '%v', age='%v'  where id='%v'", fName, lName, gender, age, id)
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
